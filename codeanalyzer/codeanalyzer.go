package main

import (
	"bufio"
	"fmt"
	"bytes"
	"strconv"
	"strings"
	"log"
	"time"
	"path"
	"os"
	"os/exec"
	"regexp"
	//	"strings"
)

type identifier struct {
	type_var  bool
	type_func bool
	count     int
	benign    bool
	vuln      bool
}

const (
    run_file_name="analyzer.out"
)

func main() {

	fl := os.Args[1]
        dr := path.Dir(os.Args[0])
	source_analyze(dr,fl)

}

func source_analyze(run_dir string, in_file string) {

	ignore_kw_map := make(map[string]bool)
	vuln_func_map := make(map[string]bool)

	kw_list := []string{"foreach", "types", "__", "wp_die", "for", "include",
		"die", "array", "unset", "isset", "if", "elseif", "defined", "define"}

	vuln_func_list := []string{"mt_rand", "eval", "file_put_contents", "assert", "base64_decode", "gzinflate"}


	var_pattern := regexp.MustCompile(`(?m)\$(?P<var>\w+)\s*(\[|\'|\w+|\'|\])*\s*\=`)
	func_pattern := regexp.MustCompile(`(?m)(?P<func>\w+)\s*\(`)
        unicode_pattern := regexp.MustCompile(`(?m)\s*\@\w+\s*\"\D*\\(?P<unicode>\d{3})`)
        // comment_pattern := regexp.MustCompile(`(?m)\s*(?P<comment>^[<//])`)


        fmt.Printf("Operating on %v\n", in_file)

	file, err := os.Open(in_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fi, err1 := file.Stat()
	if err1 != nil {
                fmt.Printf("could not open %v \n", in_file)
		return
	}

	for _, kw := range kw_list {
		ignore_kw_map[kw] = true
	}

	for _, vf := range vuln_func_list {
		vuln_func_map[vf] = true
	}

	var_funcs := make(map[string]*identifier)
	scanner := bufio.NewScanner(file)
	line_count := 0
	line_length := 0
	long_lines := 0
	max_line_length := 0
	malicious_ids := 0
        vuln_funcs := 0
        unicode_includes := 0 

	for scanner.Scan() {
		content := scanner.Text()
                line_length = len(content)
		line_count++

                if line_length > 80 {
                    long_lines++ 
                    if max_line_length < line_length {
                        max_line_length = line_length
                    }
                }

                
		fmatch := func_pattern.FindAllStringSubmatch(content, -1)
		vmatch := var_pattern.FindAllStringSubmatch(content, -1)
                umatch := unicode_pattern.FindAllStringSubmatch(content, -1)
                /*
		cmatch := comment_pattern.FindAllStringSubmatch(content, -1)

                if  len(cmatch) !=0 {
                    fmt.Println(cmatch)
                }
                */

		for i := 0; i < len(vmatch); i++ {
			if len(vmatch[i][1]) < 3 {
				continue
			}
			idp, ok := var_funcs[vmatch[i][1]]
			if !ok {
				idp := new(identifier)
				idp.type_var = true
				idp.count = 1
				idp.benign = true
				idp.vuln = false
				var_funcs[vmatch[i][1]] = idp
			} else {
				idp.type_var = true
				idp.count++
			}
		}
		for i := 0; i < len(fmatch); i++ {
			_, ok := ignore_kw_map[fmatch[i][1]]
			if ok {
				continue
			}
			idp, ok := var_funcs[fmatch[i][1]]
			if !ok {
				idp := new(identifier)
				idp.type_func = true
				idp.count = 1
				idp.benign = true
				_, vok := vuln_func_map[fmatch[i][1]]
				if vok {
					idp.vuln = true
				}
				var_funcs[fmatch[i][1]] = idp
			} else {
				idp.type_func = true
				idp.count++
			}
		}
                unicode_includes += len(umatch)
	}

        out_file_name := in_file + ".parsed"
	out_file, err := os.OpenFile(out_file_name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("cant create output file")
		os.Exit(-1)
	}

	out_file_hd := bufio.NewWriter(out_file)
	for k, _ := range var_funcs {
		fmt.Fprintf(out_file_hd, "%v\n", k)
	}
	out_file_hd.Flush()
        out_file.Close()


        detector := run_dir + "/gib_detect.py"
        fmt.Printf("about to run detector %v\n", detector)
        cmd := exec.Command("python", detector, out_file_name)

	var out bytes.Buffer

	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
            fmt.Println(err)
            log.Fatal(err)
	}

        for {
		sl, err := out.ReadString('\n')
		if err == nil {
                    s := strings.Split(strings.TrimRight(sl,"\n"), ":")
                    val, _ := strconv.ParseBool(strings.TrimSpace(s[1]))
                    v, ok := var_funcs[strings.TrimSpace(s[0])]
                    if ok {
                        (*v).benign = val
                    }
		} else {
                    break
                }
	}

        run_file, err := os.OpenFile(run_file_name , os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
            fmt.Println("Unable to write output")
            return
        }
        defer run_file.Close()
        
	var var_count, func_count int
        fmt.Fprintf(run_file,"filename \t %v Time %v \n", in_file, time.Now())

	for k, v := range var_funcs {
                var id_type, id_what, id_vuln string
                if v.type_func {
                    func_count++
                    id_type = "func"
                } else {
                    id_type = "var"
                    var_count++
                }
                if !v.benign {
                    id_what = "malicious"
                    malicious_ids += v.count
                } else {
                    id_what = "benign"
                }
                if v.vuln {
                    id_vuln = "vulnerable"
                    vuln_funcs++
		    fmt.Printf("\t %v(%v): %s %s %s\n", k, v.count,  id_vuln)
                }
                fmt.Fprintf(run_file,"\t %v(%v):\t%5.5s %10.10s %s \n", k, v.count,  id_type, id_what, id_vuln)
	}
        if var_count+func_count > 0 {
            fmt.Printf("vars = %v functions = %v lines = %v long_lines = %v \nmax_line_size = %v size = %v unicode_includes = %v malicious_ids = %v \n", 
                var_count, 
                func_count,
		line_count,
                long_lines,
                max_line_length,
                fi.Size(),
                unicode_includes,
                malicious_ids)
        }
}

/*
   results[pattern.SubexpNames()[j]] = name
*/
