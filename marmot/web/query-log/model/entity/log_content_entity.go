package entity


type LogContentEntity struct {
	Id    int64    `ddb:"id"`
	Cfrom  string `ddb:"cfrom"`
	Cto    string    `ddb:"cto"`
	Mtime string `ddb:"mtime"`
}
