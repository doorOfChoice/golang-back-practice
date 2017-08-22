<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head_write.html" .}}
    <script>
        $(function () {
            //记录要删除的标签
            $(".deleted-id").click(function (e) {
                var parent = $(this).parent()
                parent.css("display","none")
                pid = $(this).attr("data-id")
                $("#upload").append('<input type="hidden" name="deleted-ids[]" value="' + pid + '">')
            })
        })
    </script>
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}
            <div class="col-xs-10 content-back">
                <h1>修改文章</h1>
                {{if .flash.error}}
                <div class="alert alert-danger" role="alert">{{.flash.error}}</div>{{end}} {{if .flash.success}}
                <div class="alert alert-success" role="alert">{{.flash.success}}</div>{{end}} 
                {{with .Post}}
                <form id="upload" action="/admin/post/{{.Id}}" method="POST">
                    <input name="_method" value="PUT" type="hidden">
                    <div class="form-group">
                        <label for="">标题</label>
                        <input type="text" name="title" value="{{.Title}}" class="form-control .post-title">
                    </div>
                    <div class="form-group">
                        <label for="">要新增的标签</label>
                        <input type="text" name="new-tags" class="form-control .post-tags">
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">已经拥有的标签</div>
                        <div class="panel-body">
                            {{range $v := .Tags}}
                            <span class="label label-primary tag-interval">
                                {{.Name}}
                                <span class="badge tag-close deleted-id" data-id="{{.Id}}">&times;</span>
                            </span>
                            {{end}}
                        </div>
                    </div>
                    <textarea name="content" id="editable-textarea" data-provider="markdown-editable" cols="30" rows="10">{{.Content}}</textarea>
                    <div><input type="submit" name="submit" class="btn btn-primary" value="修改"></div>
                </form>
                {{end}}
            </div>
        </div>
    </div>
</body>

</html>