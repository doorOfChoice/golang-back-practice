<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head_write.html" .}}

</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}
            <div class="col-xs-10 content-back">
                <h1>发表文章</h1>
                {{if .flash.error}}
                <div class="alert alert-danger" role="alert">{{.flash.error}}</div>{{end}} {{if .flash.success}}
                <div class="alert alert-success" role="alert">{{.flash.success}}</div>{{end}}
                <form action="/admin/post" method="POST">
                    <div class="form-group">
                        <label for="">标题</label>
                        <input type="text" name="title" class=" form-control post-title" value="{{.flash.Title}}" placeholder="Title...">
                    </div>
                    <div class="form-group">
                        <label for="">标签</label>
                        <input type="text" name="tags" class="form-control post-tags" value="{{.flash.Tags}}" placeholder="Tags...">
                    </div>
                    <textarea id="editable-textarea" name="content" data-provide="markdown-editable" cols="30" rows="10">{{.flash.Content}}</textarea>
                    <div><input type="submit" name="submit" class="btn btn-primary"></div>
                </form>
            </div>
        </div>
    </div>
</body>

</html>