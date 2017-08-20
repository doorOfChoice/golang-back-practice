<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head.html" .}}
    <script>
        $(function () {
           // var deleteIds = []
            $(".deleted-id").click(function (e) {
                parent = $(this).parent()
                parent.hide()
                pid = $(this).attr("data-id")
                $("#upload").append('<input type="hidden" name="deleted-ids[]" value="' + pid + '">')
                //deleteIds.push(parseInt(pid))
            })
            $("#upload").submit(function (e) {

            })
        })
    </script>
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}
            <div class="col-xs-9">
                <h1>发表文章</h1>
                <h1 class="text-danger">{{.flash.error}}</h1>
                <h1 class="text-success">{{.flash.success}}</h1>
                {{with .Post}}
                <form id="upload" action="/admin/post/{{.Id}}" method="POST">
                    <input name="_method" value="PUT" type="hidden">
                    <ul>
                        <li><input type="text" name="title" value="{{.Title}}"></li>
                        <li>
                            {{range $v := .Tags}}
                            <span class="label label-primary">
                                {{.Name}}
                                <span class="badge deleted-id" data-id="{{.Id}}">&times;</span>
                            </span>
                            {{end}}
                        </li>
                        <li><input type="text" name="new-tags" ></li>
                        <li><textarea name="content" id="" cols="30" rows="10" >{{.Content}}</textarea></li>
                        <li><input type="submit" name="submit"></li>
                    </ul>
                </form>
                {{end}}
            </div>
        </div>
    </div>
</body>

</html>