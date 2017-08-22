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
                {{if .flash.error}}
                <div class="alert alert-danger" role="alert">{{.flash.error}}</div>{{end}} {{if .flash.success}}
                <div class="alert alert-success" role="alert">{{.flash.success}}</div>{{end}}
                <form action="/admin/user" method="POST">
                    <ul>
                        <div class="form-group">
                            <label for="">用户名</label>
                            <input type="text" name="username" class="form-control" value="{{.flash.Username}}" placeholder="username">
                        </div>
                        <div class="form-group">
                            <label for="">密码</label>
                            <input type="text" name="password" class="form-control" placeholder="password">
                        </div>
                        <div class="form-group">
                            <label for="">等级</label>
                            <select name="power" class="form-control">
                                {{range $i, $v := .Powers}}
                                    <option value="{{$i}}">{{whatpower $v}}</option>
                                {{end}}
                            </select>
                        </div>
                        <div><input class="btn btn-primary" type="submit" name="创建"></div>
                    </ul>
                </form>
            </div>
        </div>
    </div>
</body>

</html>