<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head.html" .}}
     <script>
        $(function(){
            $("#btn-query").click(function(e){
                var type = $('input[name="query-type"]:checked').val()
                var value = $("[name='query-value']").val()
                location.href = "/admin/manaUsers?type=" + type + "&value="+value
            })
        })
    </script>
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}

            <div class="col-xs-10 content-back">
                <div class="row">
                    <div class="col-xs-6">
                        <div class="btn-group">
                            <form action="/admin/user" method="GET">
                                <input type="submit" class="btn btn-primary" value="＋新建用户">
                            </form>
                        </div>
                    </div>
                    <div class="col-xs-6">
                        <div class="form-inline search-group">
                            <div class="form-group">
                                <input type="email" name="query-value" class="form-control">
                                <button class="btn btn-default" id="btn-query">查询</button>
                                <label><input type="radio" name="query-type" value="all" checked>全部</label>
                                <label><input type="radio" name="query-type" value="id">id</label>
                                <label><input type="radio" name="query-type" value="username">用户名</label>
                            </div>
                        </div>
                    </div>
                </div>
                {{if .flash.error}}<div class="alert alert-danger" role="alert">{{.flash.error}}</div>{{end}}
                {{if .flash.success}}<div class="alert alert-success" role="alert">{{.flash.success}}</div>{{end}}
                <table class="table table-hover">
                    <tr>
                        <th>id</th>
                        <th>用户名</th>
                        <th>权限</th>
                        <th>操作</th>
                    </tr>
                    {{range $v := .Users}} {{with $v}}
                    <tr>
                        <td>{{.Id}}</td>
                        <td>{{.Username}}</td>
                        <td>{{whatpower .Identification}}</td>
                        <td>
                            <form action="/admin/user/{{.Id}}" method="POST">
                                <input type="hidden" name="_method" value="DELETE">
                                <input type="submit" class="btn btn-danger" value="删除">
                            </form>
                            <form action="/admin/user/{{.Id}}" method="GET">
                                <input type="submit" class="btn btn-success" value="修改">
                            </form>
                        </td>
                    </tr>
                    {{end}} {{end}}
                </table>
                <div>
                    <span>页码 {{.P.CurrentPage}}/{{.P.MaxPage}},</span>
                    <span>显示 {{.P.CurrentValue}}/{{.P.PerValue}}, 总共检索到{{.P.MaxValue}}条数据</span>
                    <nav aria-label="Page navigation">
                        <ul class="pagination">
                            <li><a href="{{.P.PageLinkFirst}}"><span aria-hidden="true">首页</span></a></li>
                            <li><a href="{{.P.PageLinkPrev}}"><span aria-hidden="true">&laquo;</span></a></li>
                            {{range $v := .P.Links}} {{with $v}}
                            <li {{if .IsCurrent}}class="active" {{end}}>
                                <a href="{{.Href}}">{{.Id}}</a>
                            </li>
                            {{end}} {{end}}
                            <li><a href="{{.P.PageLinkNext}}"><span aria-hidden="true">&raquo;</span></a></li>
                            <li><a href="{{.P.PageLinkLast}}"><span aria-hidden="true">尾页</span></a></li>
                        </ul>
                    </nav>
                </div>
            </div>
        </div>
    </div>
</body>

</html>