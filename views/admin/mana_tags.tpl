<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head.html" .}}
    <script>
        $(function () {
            var panelTagAdd = $("#panel-add-tag")
            $("#show-add").click(function (e) {
                panelTagAdd.fadeToggle(500)
            })
        })
    </script>
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}

            <div class="col-xs-10 content-back">
                <div>
                    <div class="row">
                        <div class="col-xs-6">
                            <div class="btn-group">
                                <button class="btn btn-primary" id="show-add">＋新建标签</button>
                            </div>
                        </div>
                        <div class="col-xs-6">
                            <div class="dropdown search-group">
                                <button class="btn btn-primary dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true"
                                    aria-expanded="true">
   排序
    <span class="caret"></span>
  </button>
                                <ul class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenu1">
                                    <li><a href="{{urlfor `TagController.ManaPage`}}?order=desc">文章降序</a></li>
                                    <li><a href="{{urlfor `TagController.ManaPage`}}?order=asc">文章升序</a></li>
                                    <li role="separator" class="divider"></li>
                                    <li><a href="{{urlfor `TagController.ManaPage`}}?order=id_asc">id升序</a></li>
                                    <li><a href="{{urlfor `TagController.ManaPage`}}?order=id_desc">id降序</a></li>
                                </ul>
                        </div>
                        <!-- <div class="form-inline search-group">
                                
                                <div class="form-group">
                                    <button class="btn btn-primary" onclick="location.href='{{urlfor `TagController.ManaPage`}}?order=asc'">升序</button>
                                    <button class="btn btn-primary" onclick="location.href='{{urlfor `TagController.ManaPage`}}?order=desc'">降序</button>
                                </div>
                            </div> -->
                    </div>
                </div>

                <div id="panel-add-tag" class="btn-group" style="display:none">
                    <form action="/admin/tag" method="POST">
                        <div class="input-group">
                            <input type="text" class="form-control" name="name" placeholder="Add to...">
                            <span class="input-group-btn">
                                    <input class="btn btn-default" type="submit" value="添加">
                                </span>
                        </div>
                    </form>
                </div>
            </div>
            {{if .flash.error}}<div class="alert alert-danger" role="alert">{{.flash.error}}</div>{{end}}
            {{if .flash.success}}<div class="alert alert-success" role="alert">{{.flash.success}}</div>{{end}}
            <table class="table table-hover">
                <tr>
                    <th>id</th>
                    <th>标签名</th>
                    <th>拥有文章</th>
                    <th>操作</th>
                </tr>
                {{range $v := .Tags}} {{with $v}}
                <tr>
                    <td>{{.Id}}</td>
                    <td>{{.Name}}</td>
                    <td><a href="{{urlfor `PostController.ManaPage`}}?type=tag&value={{.Name}}">{{len .Posts}}</a></td>
                    <td>
                        <form action="/admin/tag/{{.Id}}" method="POST">
                            <input type="hidden" name="_method" value="DELETE">
                            <input type="submit" class="btn btn-danger" value="删除">
                        </form>
                    </td>
                </tr>
                {{end}} {{end}}
            </table>
            {{$len := len .Tags}}
                {{if eq $len 0}}
                    <center><h3>😄空空如也...</h3></center>
                {{end}}
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