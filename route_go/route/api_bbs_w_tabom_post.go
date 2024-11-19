package route

import (
    "database/sql"
    "opennamu/route/tool"
    
    jsoniter "github.com/json-iterator/go"
)

func Api_bbs_w_post_tabom_post(db *sql.DB, call_arg []string) string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

    other_set := map[string]string{}
    json.Unmarshal([]byte(call_arg[0]), &other_set)

    return_data := make(map[string]interface{})

    if tool.Check_acl(db, "", "", "bbs_comment", other_set["ip"]) {
        return_data["response"] = "require auth"

        json_data, _ := json.Marshal(return_data)
        return string(json_data)
    }

	return ""
}