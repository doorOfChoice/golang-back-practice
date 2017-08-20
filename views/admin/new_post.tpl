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
                <h1>发表文章</h1>
                <p class="text-danger">{{.flash.error}}</p>
                <p class="text-success">{{.flash.success}}</p>
                <form action="/admin/post" method="POST">
                    <ul>
                        <li><input type="text" name="title" value="{{.flash.Title}}"></li>
                        <li><input type="text" name="tags" value="{{.flash.Tags}}"></li>
                        <li><textarea name="content" id="" cols="30" rows="10" >{{.flash.Content}}</textarea></li>
                        <li><input type="submit" name="submit"></li>
                    </ul>
                </form>
            </div>
        </div>
    </div>
</body>

</html>