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
                <h1>创建用户</h1>
                <p class="text-danger">{{.flash.error}}</p>
                <p class="text-success">{{.flash.success}}</p>
                <form action="/admin/user" method="POST">
                    <ul>
                        <li><input type="text" name="username" value="{{.flash.Username}}" placeholder="username"></li>
                        <li><input type="text" name="password" placeholder="password"></li>
                        <li>
                            <select name="power">
                                {{range $i, $v := .Powers}}
                                    <option value="{{$i}}">{{whatpower $v}}</option>
                                {{end}}
                            </select>
                        </li>
                        <li><input type="submit" name="创建"></li>
                    </ul>
                </form>
            </div>
        </div>
    </div>
</body>

</html>