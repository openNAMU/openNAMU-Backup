package main

import (
    "os"
    "log"
    "io/ioutil"
    "strings"
    "sync"
    "runtime"
    "opennamu/route"
    "opennamu/route/tool"
    
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
        
    if len(os.Args) > 1 && os.Args[1] == "dev" {
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

    tool.DB_init()

    db_type := tool.Get_DB_type()
    if db_type == "sqlite" {
        runtime.GOMAXPROCS(1)
    }

    var mu sync.Mutex
    
    r := gin.Default()
    r.POST("/", func(c *gin.Context) {
        route_data := ""
        body, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            log.Fatal(err)
        }
        
        body_string := string(body)
        word := strings.Fields(body_string)
        
        call_arg := []string{ word[0], strings.Join(word[1:], " ") }

        if db_type == "sqlite" {
            mu.Lock()
            defer mu.Unlock()
        }

        db := tool.DB_connect()
        defer tool.DB_close(db)
        
        if call_arg[0] == "test" {
            route_data = "ok"
        } else if call_arg[0] == "main_func_easter_egg" {
            route_data = route.Main_func_easter_egg()
        } else if call_arg[0] == "api_w_raw" {
            route_data = route.Api_w_raw(db, call_arg[1:])
        } else if call_arg[0] == "api_func_sha224" {
            route_data = route.Api_func_sha224(db, call_arg[1:])
        } else if call_arg[0] == "api_w_random" {
            route_data = route.Api_w_random(db, call_arg[1:])
        } else if call_arg[0] == "api_func_search" {
            route_data = route.Api_func_search(db, call_arg[1:])
        } else if call_arg[0] == "api_topic" {
            route_data = route.Api_topic(db, call_arg[1:])
        } else if call_arg[0] == "api_func_ip" {
            route_data = route.Api_func_ip(db, call_arg[1:])
        } else if call_arg[0] == "api_list_recent_change" {
            route_data = route.Api_list_recent_change(db, call_arg[1:])
        } else if call_arg[0] == "api_list_recent_edit_request" {
            route_data = route.Api_list_recent_edit_request(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs" {
            route_data = route.Api_bbs(db, call_arg[1:])
        } else if call_arg[0] == "api_w_xref" {
            route_data = route.Api_w_xref(db, call_arg[1:])
        } else if call_arg[0] == "api_w_watch_list" {
            route_data = route.Api_w_watch_list(db, call_arg[1:])
        } else if call_arg[0] == "api_user_watch_list" {
            route_data = route.Api_user_watch_list(db, call_arg[1:])
        } else if call_arg[0] == "api_w_render" {
            route_data = route.Api_w_render(db, call_arg[1:])
        } else if call_arg[0] == "api_func_llm" {
            route_data = route.Api_func_llm(db, call_arg[1:])
        } else if call_arg[0] == "api_func_language" {
            route_data = route.Api_func_language(db, call_arg[1:])
        } else if call_arg[0] == "api_func_auth" {
            route_data = route.Api_func_auth(db, call_arg[1:])
        } else if call_arg[0] == "api_list_recent_discuss" {
            route_data = route.Api_list_recent_discuss(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_list" {
            route_data = route.Api_bbs_list(db, call_arg[1:])
        } else if call_arg[0] == "api_list_old_page" {
            route_data = route.Api_list_old_page(db, call_arg[1:])
        } else if call_arg[0] == "api_topic_list" {
            route_data = route.Api_topic_list(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w_n" {
            route_data = route.Api_bbs_w(db, call_arg[1:])
        } else if call_arg[0] == "api_w_set_reset" {
            route_data = route.Api_w_set_reset(db, call_arg[1:])
        } else if call_arg[0] == "api_list_recent_block" {
            route_data = route.Api_list_recent_block(db, call_arg[1:])
        } else if call_arg[0] == "api_list_title_index" {
            route_data = route.Api_list_title_index(db, call_arg[1:])
        } else if call_arg[0] == "api_user_setting_editor_post" {
            route_data = route.Api_user_setting_editor_post(db, call_arg[1:])
        } else if call_arg[0] == "api_user_setting_editor_delete" {
            route_data = route.Api_user_setting_editor_delete(db, call_arg[1:])
        } else if call_arg[0] == "api_user_setting_editor" {
            route_data = route.Api_user_setting_editor(db, call_arg[1:])
        } else if call_arg[0] == "api_setting" {
            route_data = route.Api_setting(db, call_arg[1:])
        } else if call_arg[0] == "api_setting_put" {
            route_data = route.Api_setting_put(db, call_arg[1:])
        } else if call_arg[0] == "api_func_ip_menu" {
            route_data = route.Api_func_ip_menu(db, call_arg[1:])
        } else if call_arg[0] == "api_func_ip_post" {
            route_data = route.Api_func_ip_post(db, call_arg[1:])
        } else if call_arg[0] == "api_list_acl" {
            route_data = route.Api_list_acl(db, call_arg[1:])
        } else if call_arg[0] == "api_user_rankup" {
            route_data = route.Api_user_rankup(db, call_arg[1:])
        } else if call_arg[0] == "api_func_acl" {
            route_data = route.Api_func_acl(db, call_arg[1:])
        } else if call_arg[0] == "api_func_ban" {
            route_data = route.Api_func_ban(db, call_arg[1:])
        } else if call_arg[0] == "api_func_auth_post" {
            route_data = route.Api_func_auth_post(db, call_arg[1:])
        } else if call_arg[0] == "api_give_auth_patch" {
            route_data = route.Api_give_auth_patch(db, call_arg[1:])
        } else if call_arg[0] == "api_list_auth" {
            route_data = route.Api_list_auth(db, call_arg[1:])
        } else if call_arg[0] == "api_w_page_view" {
            route_data = route.Api_w_page_view(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w_comment_one" {
            route_data = route.Api_bbs_w_comment_one(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w_comment" {
            route_data = route.Api_bbs_w_comment(db, call_arg[1:])
        } else if call_arg[0] == "api_list_history" {
            route_data = route.Api_list_history(db, call_arg[1:])
        } else if call_arg[0] == "api_list_markup" {
            route_data = route.Api_list_markup(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w_set" {
            route_data = route.Api_bbs_w_set(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w_set_put" {
            route_data = route.Api_bbs_w_set_put(db, call_arg[1:])
        } else if call_arg[0] == "api_func_alarm_post" {
            route_data = route.Api_func_alarm_post(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w" {
            route_data = route.Api_bbs_w(db, call_arg[1:])
        } else if call_arg[0] == "api_bbs_w_post" {
            route_data = route.Api_bbs_w_post(db, call_arg[1:])
        } else if call_arg[0] == "api_w_comment" {
            route_data = route.Api_w_comment(db, call_arg[1:])
        }
    
        c.String(http.StatusOK, route_data)  
    })
    
    r.Run(":" + tool.Get_port())
}
