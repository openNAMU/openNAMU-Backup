from .tool.func import *

def list_please_2(conn):
    curs = conn.cursor()

    num = int(number_check(flask.request.args.get('num', '1')))
    if num * 50 > 0:
        sql_num = num * 50 - 50
    else:
        sql_num = 0

    div = '<ul>'

    curs.execute(db_change("select distinct title from back where type = 'no' order by title asc limit ?, 50"), [sql_num])
    data_list = curs.fetchall()
    for data in data_list:
        div += '<li><a class="not_thing" href="/w/' + url_pas(data[0]) + '">' + data[0] + '</a></li>'

    div += '</ul>' + next_fix('/please?num=', num, data_list)

    return easy_minify(flask.render_template(skin_check(),
        imp = [load_lang('need_document'), wiki_set(), custom(), other2([0, 0])],
        data = div,
        menu = [['other', load_lang('return')]]
    ))