<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head.html" .}}
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}

            <div class="col-xs-10 content-back">
                <div class="btn-group" >
                    <form action="/admin/post" method="GET">
                        <input type="submit" class="btn btn-primary" value="＋新建文章">
                    </form>
                </div>
                <p class="text-danger">{{.flash.error}}</p>
                <p class="text-success">{{.flash.success}}</p>
                <table class="table table-hover">
                    <tr>
                        <th>id</th>
                        <th>标题</th>
                        <th>创建日期</th>
                        <th>修改日期</th>
                        <th>行为</th>
                    </tr>
                    {{range $v := .Posts}} {{with $v}}
                    <tr>
                        <td>{{.Id}}</td>
                        <td>{{.Title}}</td>
                        <td>{{date .CreateDate "Y-m-d H:i:s"}}</td>
                        <td>{{date .UpdateDate "Y-m-d H:i:s"}}</td>
                        <td>
                            <form action="/admin/post/{{.Id}}" method="POST">
                                <input type="hidden" name="_method" value="DELETE">
                                <input type="submit" class="btn btn-danger" value="删除">
                            </form>
                            <form action="/admin/post/{{.Id}}" method="GET">
                                <input type="submit" class="btn btn-success" value="查看">
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