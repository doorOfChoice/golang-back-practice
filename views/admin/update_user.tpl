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
                <h1>修改用户</h1>
                <p class="text-danger">{{.flash.error}}</p>
                <p class="text-success">{{.flash.success}}</p>
                <form action="/admin/user/{{.User.Id}}" method="POST">
                    <input name="_method" value="PUT" type="hidden">
                    <ul>
                        <li><label for="">用户名:</label>{{.User.Username}}</li>
                        <li><input type="text" name="password" placeholder="password"></li>
                        <li>
                            <select name="power" value="{{whatpower .User.Identification}}">
                                {{range $i, $v := .Powers}}
                                <option value="{{$i}}" {{if eq $.User.Identification $v}}selected{{end}}>
                                    {{whatpower $v}}
                                </option>
                                {{end}}
                            </select>
                        </li>
                        <li><input type="submit" name="修改"></li>
                    </ul>
                </form>
            </div>
        </div>
    </div>
</body>

</html>