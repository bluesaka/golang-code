/**
@link https://goframe.org/display/gf/gen+model

$ gf gen model -h
USAGE
    gf gen model [OPTION]

OPTION
    -/--path             directory path for generated files.
    -l, --link           database configuration, the same as the ORM configuration of GoFrame.
    -t, --tables         generate models only for given tables, multiple table names separated with ','
    -g, --group          specifying the configuration group name for database,
                         it's not necessary and the default value is "default"
    -c, --config         used to specify the configuration file for database, it's commonly not necessary.
                         If "-l" is not passed, it will search "./config.toml" and "./config/config.toml"
                         in current working directory in default.
    -p, --prefix         add prefix for all table of specified link/database tables.
    -r, --removePrefix   remove specified prefix of the table, multiple prefix separated with ','
    -m, --mod            module name for generated golang file imports.

CONFIGURATION SUPPORT
    Options are also supported by configuration file. The configuration node name is "gf.gen", which also supports
    multiple databases, for example:
    [gfcli]
        [[gfcli.gen.model]]
            link   = "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
            tables = "order,products"
        [[gfcli.gen.model]]
            link   = "mysql:root:12345678@tcp(127.0.0.1:3306)/primary"
            path   = "./my-app"
            prefix = "primary_"
            tables = "user, userDetail"

EXAMPLES
    gf gen model
    gf gen model -l "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
    gf gen model -path ./model -c config.yaml -g user-center -t user,user_detail,user_login
    gf gen model -r user_

DESCRIPTION
    The "gen" command is designed for multiple generating purposes.
    It's currently supporting generating go files for ORM models.


install gf-cli
```
wget https://goframe.org/cli/darwin_amd64/gf && chmod +x gf && ./gf install
If you're using zsh, you might need rename your alias by command alias gf=gf to resolve the conflicts between gf and git fetch.
then exec `source ~/.zshrc` to use
```

// need go.mod
gf gen dao -l "mysql:root:123456@tcp(127.0.0.1:3306)/test" -t user,product -path ./app/model

gf gen model -l "mysql:root:123456@tcp(127.0.0.1:3306)/test" -t user,product -path ./app/model

 */
package model
