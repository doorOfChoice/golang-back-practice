<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head.html" .}}
    <style>
        .operation a{
            font-size: 20px;
        }
    </style>
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}} {{with .LoginUser}}
            <div class="col-xs-10 content-back">
                <div>
                    <div class="col-md-6">
                        <div class="panel panel-default">
                            <div class="panel-heading">个人资料</div>
                            <div class="panel-body">
                                <div class="personal-image">
                                    <div class="col-md-4">
                                        <center><img src="{{.Profile.Head}}" class="img-header" alt=""></center>
                                    </div>
                                    <div class="col-md-8">
                                        <p>用户名: {{.Username}}</p>
                                        <p>权限: {{whatpower .Identification}}</p>
                                        <p>姓名: {{.Profile.Name}}</p>
                                        <p>爱好: {{.Profile.Hobby}}</p>

                                        <div class="operation">
                                            <form name="form" action="/admin/user" method="POST">
                                                <input type="hidden" name="_method" value="DELETE">
                                                <a href="javascript:document.form.submit()"><span class="glyphicon glyphicon-log-out"></span></a>
                                            </form>
                                            <a href="/admin/user/{{.Id}}"><span class="glyphicon glyphicon-user"></span></a>
                                            <a href="/admin/post"><span class="glyphicon glyphicon-pencil"></span></a>
                                        </div>
                                    </div>
                                </div>

                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-6">
                    <div class="panel panel-default">
                        <div class="panel-heading">个人站点</div>
                        <div class="panel-body">
                            <ul class="list-group">
                                {{$vs := stations .Profile.Stations}} {{range $vs}}
                                <li class="list-group-item">
                                    <img src="{{.Icon}}" alt="">
                                    <span><b>{{.Name}}</b></span>
                                    <p><a href="{{.Href}}">{{.Href}}</a></p>
                                </li>
                                {{end}}
                            </ul>
                        </div>
                    </div>
                </div>
            </div>

        </div>
        {{end}}
    </div>
    </div>
</body>

</html>