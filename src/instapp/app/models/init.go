package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "gopkg.in/gorp.v1"
    "log"
    elasticapi "gopkg.in/olivere/elastic.v7"
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    loghttp "github.com/motemen/go-loghttp"
    "bytes"
    "time"
    "errors"
    "github.com/motemen/go-nuts/roundtime"

)


// Connection to database
// TODO: get from env
const (
	DBHOST = "tcp(mysql:3306)"
    // DBHOST = "tcp(mysql:3306)"
	DBNAME = "instabug"
	DBUSER = "root"
	DBPASS = "123"
    ELASTICSEARCH_URI = "http://elasticsearch:9200"
)

// Global database references
var db *sql.DB
var dbmap *gorp.DbMap

func InitDB(){
	var err error

    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err = sql.Open("mysql", DBUSER+":"+DBPASS+"@"+DBHOST+"/"+DBNAME)

    // construct a gorp DbMap
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    logErrDb(err, "sql.Open failed")

    // create the table. in a production system you'd generally
    // use a migration tool, or create the tables via scripts
    buildDatabaseTables()
    err = dbmap.CreateTablesIfNotExists()
    logErrDb(err, "Create tables failed")
}

// Build tables
func buildDatabaseTables() {
	dbmap.AddTableWithName(Application{}, "applications").SetKeys(true, "Id")
    dbmap.AddTableWithName(Chat{}, "chats").SetKeys(true, "Id")
    dbmap.AddTableWithName(Message{}, "messages").SetKeys(true, "Id")
}

// log database errors to console
func logErrDb(err error, msg string) {
    if err != nil {
        log.Fatalln("DB ERR", err)
    }
}

func InitElastic(ctx context.Context, url string, sniff bool, responseSize int) (*elasticapi.Client, error) {

    var httpClient = &http.Client{
        Transport: &loghttp.Transport{
            LogRequest: func(req *http.Request) {
                var bodyBuffer []byte
                if req.Body != nil {
                    bodyBuffer, _ = ioutil.ReadAll(req.Body) // after this operation body will equal 0
                    // Restore the io.ReadCloser to request
                    req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))
                }
                fmt.Println("--------- Elasticsearch ---------")
                fmt.Println("Request URL : ", req.URL)
                fmt.Println("Request Method : ", req.Method)
                fmt.Println("Request Body : ", string(bodyBuffer))
            },
            LogResponse: func(resp *http.Response) {
                ctx := resp.Request.Context()
                if start, ok := ctx.Value(loghttp.ContextKeyRequestStart).(time.Time); ok {
                    fmt.Println("Response Status : ", resp.StatusCode)
                    fmt.Println("Response Duration : ", roundtime.Duration(time.Now().Sub(start), 2))
                } else {
                    fmt.Println("Response Status : ", resp.StatusCode)
                }
                fmt.Println("--------------------------------")
            },
        },
    }

    client, err := elasticapi.NewClient(elasticapi.SetURL(url), elasticapi.SetSniff(sniff), elasticapi.SetHttpClient(httpClient))
    if err != nil {
        return nil, err
    }

    err = ping(ctx, client, url)
    if err != nil {
        return nil, err
    }

    return client, nil
}

// Ping method
func ping(ctx context.Context, client *elasticapi.Client, url string) error {

    // Ping the Elasticsearch server to get HttpStatus, version number
    if client != nil {
        info, code, err := client.Ping(url).Do(ctx)
        if err != nil {
            return err
        }

        fmt.Printf("Elasticsearch returned with code %d and version %s \n", code, info.Version.Number)
        return nil
    }
    return errors.New("elastic client is nil")
}



