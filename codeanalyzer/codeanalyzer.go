package main

import (
	"bufio"
	"fmt"
	"bytes"
	"strconv"
	"strings"
	"log"
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

func main() {

	fl := os.Args[1]
	source_analyze(fl)

}

func source_analyze(in_file string) {

	ignore_kw_map := make(map[string]bool)
	vuln_func_map := make(map[string]bool)

	kw_list := []string{"foreach", "types", "__", "wp_die", "for", "include",
		"die", "array", "unset", "isset", "if", "elseif", "defined", "define"}
	vuln_func_list := []string{"mt_rand", "eval", "file_put_contents", "assert", "base64_decode", "gzinflate"}
	//var_pattern := regexp.MustCompile(`(?m)^[^<]*\$(?P<var>\w+)\s*\=`)
	var_pattern := regexp.MustCompile(`(?m)\$(?P<var>\w+)\s*(\[|\'|\w+|\'|\])*\s*\=`)
	func_pattern := regexp.MustCompile(`(?m)(?P<func>\w+)\s*\(`)
        // comment_pattern := regexp.MustCompile(`(?m)\s*(?P<comment>^[<//])`)

	fmt.Printf("File %v \n", in_file)

	file, err := os.Open(in_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fi, err1 := file.Stat()
	if err1 != nil {
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

	for scanner.Scan() {
		content := scanner.Text()
		line_count++

                
		fmatch := func_pattern.FindAllStringSubmatch(content, -1)
		vmatch := var_pattern.FindAllStringSubmatch(content, -1)
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
	}

        out_file_name := in_file + ".parsed"
	out_file, err := os.OpenFile(out_file_name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("cant create output file")
		os.Exit(1)
	}

	out_file_hd := bufio.NewWriter(out_file)
	for k, _ := range var_funcs {
		fmt.Fprintf(out_file_hd, "%v\n", k)
	}
	out_file_hd.Flush()
        out_file.Close()

        cmd := exec.Command("python", "gib_detect.py", out_file_name)

	var out bytes.Buffer

	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
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

	var var_count, func_count int
	for k, v := range var_funcs {
                var id_type, id_what, id_vuln string
                if v.type_func {
                    id_type = "func"
                } else {
                    id_type = "var"
                }
                if !v.benign {
                    id_what = "malicious"
                } else {
                    id_what = "benign"
                }
                if v.vuln {
                    id_vuln = "vulnerable"
                }
		fmt.Printf("\t %v(%v):,%s %s %s\n", k, v.count,  id_type,
			id_what, id_vuln)
		if v.type_var {
			var_count++
		}
		if v.type_func {
			func_count++
		}
	}
        if var_count+func_count > 0 {
	fmt.Printf("variables %v function_identifiers %v Total %v  lines %v  size %v \n", var_count, func_count,
		var_count+func_count, line_count, fi.Size())
        }
}

/*
   results[pattern.SubexpNames()[j]] = name
*/
