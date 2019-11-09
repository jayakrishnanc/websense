<?php
if ( !defined('WP_ADMIN') )
	define('WP_ADMIN', TRUE);
if ( isset($_GET['import']) && !defined('WP_LOAD_IMPORTERS') )
	define('WP_LOAD_IMPORTERS', true);
require_once(dirname(dirname(__FILE__)) . '/wp-load.php');
if ( get_option('db_upgraded') ) {
	$wp_rewrite->flush_rules();
	update_option( 'db_upgraded',  false );
	do_action('after_db_upgrade');
} elseif ( get_option('db_version') != $wp_db_version ) {
	if ( !is_multisite() ) {
		wp_redirect(admin_url('upgrade.php?_wp_http_referer=' . urlencode(stripslashes($_SERVER['REQUEST_URI']))));
		exit;
	} elseif ( apply_filters( 'do_mu_upgrade', true ) ) {
		$c = get_blog_count();
		if ( $c <= 50 || ( $c > 50 && mt_rand( 0, (int)( $c / 50 ) ) == 1 ) ) {
			require_once( ABSPATH . WPINC . '/http.php' );
			$response = wp_remote_get( admin_url( 'upgrade.php?step=1' ), array( 'timeout' => 120, 'httpversion' => '1.1' ) );
			do_action( 'after_mu_upgrade', $response );
			unset($response);
		}
		unset($c);
	}
}
require_once(ABSPATH . 'wp-admin/includes/admin.php');
auth_redirect();
nocache_headers();
update_category_cache();
// Schedule trash collection
if ( !wp_next_scheduled('wp_scheduled_delete') && !defined('WP_INSTALLING') )
	wp_schedule_event(time(), 'daily', 'wp_scheduled_delete');
set_screen_options();
$date_format = get_option('date_format');
$time_format = get_option('time_format');
wp_reset_vars(array('profile', 'redirect', 'redirect_url', 'a', 'text', 'trackback', 'pingback'));
wp_enqueue_script( 'common' );
wp_enqueue_script( 'jquery-color' );
$editing = false;
if ( isset($_GET['page']) ) {
	$plugin_page = stripslashes($_GET['page']);
	$plugin_page = plugin_basename($plugin_page);
}
if ( isset($_GET['post_type']) )
	$typenow = sanitize_key($_GET['post_type']);
else
	$typenow = '';
if ( isset($_GET['taxonomy']) )
	$taxnow = sanitize_key($_GET['taxonomy']);
else
	$taxnow = '';
