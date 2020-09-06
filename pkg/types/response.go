package types

/*
{
"results" :
[
{
  "repo" : "jcenter-mvn-remote-cache",
  "path" : "org/apache/maven/plugins/maven-jar-plugin/2.4",
  "name" : "maven-jar-plugin-2.4.pom",
  "type" : "file",
  "size" : 5834,
  "created" : "2020-07-21T21:08:08.273Z",
  "created_by" : "admin",
  "modified" : "2012-01-26T18:02:54.000Z",
  "modified_by" : "admin",
  "updated" : "2020-07-21T21:08:08.281Z",
  "stats" : [ {
    "downloaded" : "2020-07-21T21:08:15.651Z",
    "downloaded_by" : "token:jfxidx@01dytespze20z1ar78bvd3fbnz",
    "downloads" : 3,
    "remote_downloads" : 0
  }]
},
{
  "repo" : "jcenter-mvn-remote-cache",
  "path" : "org/apache/maven/plugins/maven-plugins/22",
  "name" : "maven-plugins-22.pom",
  "type" : "file",
  "size" : 13039,
  "created" : "2020-07-21T21:08:09.627Z",
  "created_by" : "admin",
  "modified" : "2011-08-15T20:29:16.000Z",
  "modified_by" : "admin",
  "updated" : "2020-07-21T21:08:09.629Z",
  "stats" : [ {
    "downloaded" : "2020-07-21T21:08:10.111Z",
    "downloaded_by" : "admin",
    "downloads" : 2,
    "remote_downloads" : 0
  } ]
} ],
"range" : {
  "start_pos" : 0,
  "end_pos" : 2,
  "total" : 2,
  "limit" : 2
}
}

*/

type Artifacts struct {
	Results []Results `json:"results"`
}

type Results struct {
	Repo       string  `json:"repo"`
	Path       string  `json:"path"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Size       int     `json:"size"`
	Created    string  `json:"created"`
	CreatedBy  string  `json:"created_by"`
	Modified   string  `json:"modified"`
	ModifiedBy string  `json:"modified_by"`
	Updated    string  `json:"updated"`
	Stats      []Stats `json:"stats",omitempty`
	Range      Range   `json:"range"`
}

type Stats struct {
	Downloaded      string `json:"downloaded"`
	DownloadedBy    string `json:"downloaded_by"`
	Downloads       int    `json:"downloads"`
	RemoteDownloads int    `json:"remote_downloads"`
}

type Range struct {
	StartPos int `json:"start_pos"`
	EndPos   int `json:"end_pos"`
	Total    int `json:"total"`
	Limit    int `json:"limit"`
}
