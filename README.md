# tiny webdav server (twds)

`twds` can be started without command line parameters

	twds

After starting it it will output access and error log messages to <stderr>.  
`twds` stops its operation by typing `CTRL-C`.

However before we can start using `twds` it has to be configured.

## Configuring twds.

There has to be a config file in json format in `"${USER_CONFIG_DIR}/twds/config.json"`.
`${USER_CONFIG_DIR}` is usually set to `"${HOME}/.config"` under linux.
The following json file shows an example with the default assumptions for
`listen`, `prefix`, `space`.  Edit this file if you want other settings.  Remove
a line if you want to retain the default assumptions.

	{
	   "listen": "127.0.0.0:8080",
	   "prefix": "/",
	   "space": ".twds"
	}

Parameter `listen` sets the socket for binding the server.  `prefix` sets the prefix for
each webdav item in its URL.  Parameter `space` sets a directory on which `twds` operates
on.  Here the files are stored which twds transmit to client or client sent to `twds`.
If `space` contains a relative file name this file name relates to user's home directory.
