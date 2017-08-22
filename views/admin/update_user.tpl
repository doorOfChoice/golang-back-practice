<!DOCTYPE html>
<html>

<head>
    {{template "admin/layout/head.html" .}}
    <script>
        $(function () {
            var fileHead = $("#file-head");
            var headView = $("#view-head");
            var ok = $("#ok-head")
            var fail = $("#fail-head")
            var net = $("#upload-net-head")
            $("#upload-local-head").click(function (e) {
                e.preventDefault()
                fileHead.click()
                fileHead.change(function (e) {
                    if (fileHead[0] && fileHead[0].files[0]) {
                        var reader = new FileReader();
                        reader.onload = function (e) { headView.attr("src", e.target.result); }
                        reader.readAsDataURL(fileHead[0].files[0])
                    }
                })
            })
            $("#upload-net-head").click(function (e) {
                e.preventDefault()
                $(this).attr("disabled", "")
                $.ajax({
                    type: "POST",
                    url: "https://sm.ms/api/upload",
                    processData: false,
                    contentType: false,
                    data: new FormData($("#form-head")[0])
                }).done(function (e) {
                    ok.css("display", "none")
                    fail.css("display", "none")
                    if (e.code == "success") {
                        ok.text("设置成功,点击修改上传");
                        ok.css("display", "block")
                        $("#user-head").val(e.data.url);
                    } else {
                        fail.css("display", "block")
                        fail.text("错误:" + e.msg);
                    }
                    net.removeAttr("disabled")
                })
            })
        })
    </script>
</head>

<body>
    <div class="container-fluid vertical-block">
        <div class="row vertical-block">
            {{template "admin/layout/navigator.html" .}}
            <div class="col-xs-10 content-back">
                {{if .flash.error}}
                <div class="alert alert-danger" role="alert">{{.flash.error}}</div>{{end}} {{if .flash.success}}
                <div class="alert alert-success" role="alert">{{.flash.success}}</div>{{end}}
                <div class="panel panel-primary">
                    <div class="panel-heading">
                        <h1>修改账户信息</h1>
                    </div>
                    <div class="panel-body">
                        <form action="/admin/user/{{.User.Id}}" method="POST">
                            <input name="_method" value="PUT" type="hidden">
                            <div class="form-group">
                                <label for="">用户名</label>
                                <input type="text" name="username" class="form-control" value="{{.User.Username}}" placeholder="username" disabled>
                            </div>
                            <div class="form-group">
                                <label for="">密码</label>
                                <input type="text" name="password" class="form-control" placeholder="password">
                            </div>
                            {{if not .IsSelf}}
                            <div class="form-group">
                                <label for="">等级</label>
                                <select name="power form-control">
                                {{range $i, $v := .Powers}}
                                    <option value="{{$i}}"
                                    {{if eq $.User.Identification $i}}
                                        selected
                                    {{end}}
                                    >{{whatpower $v}}</option>
                                {{end}}
                            </select>
                            </div>
                            {{end}}
                            <div><input class="btn btn-primary" type="submit" name="创建"></div>
                        </form>
                    </div>
                </div>


                {{if .IsSelf}}
                <div class="panel panel-primary">
                    <div class="panel-heading">
                        <h1>修改个人信息</h1>
                    </div>
                    {{with .User.Profile}}
                    <div class="panel-body">
                        <div class="col-md-4">
                            <center>
                                <img id="view-head" src="{{.Head}}" class="img-header" alt="">
                                <p id="ok-head" class="text-success"></p>
                                <p id="fail-head" class="text-danger"></p>
                                <div class="btn-group">
                                    <form id="form-head" action="https://sm.ms/api/upload" enctype="multipart/form-data">
                                        <button id="upload-local-head" class="btn btn-primary">预览头像</button>
                                        <button id="upload-net-head" class="btn btn-success">设定头像</button>
                                        <input id="file-head" name="smfile" type="file" style="display:none">
                                    </form>
                                </div>
                            </center>
                        </div>
                        <form action="/admin/profile/{{.Id}}" method="POST">
                            <input name="_method" value="PUT" type="hidden">
                            <input type="hidden" id="user-head" name="user-head" value="{{.Head}}">
                            <div class="col-md-8">
                                <div class="form-group">
                                    <label for="">姓名</label>
                                    <input type="text" class="form-control" name="user-name" value="{{.Name}}" placeholder="姓名">
                                </div>
                                <div class="form-group">
                                    <label for="">爱好</label>
                                    <input type="text" class="form-control" name="user-hobby" value="{{.Hobby}}" placeholder="爱好...">
                                </div>
                                <div class="form-group">
                                    <label for="">站点</label>
                                    <input type="text" class="form-control" name="user-stations" value="{{.Stations}}" placeholder="站点...">
                                </div>
                                <div class="form-group">
                                    <label for="">说说</label>
                                    <textarea type="text" class="form-control" name="user-introduction" placeholder="说点什么吧..." style="height:150px;">{{.Introduction}}</textarea>
                                </div>
                                <input type="submit" class="btn btn-primary" value="修改">
                        </form>
                        </div>
                    </div>
                    {{end}}
                </div>
                {{end}}
            </div>
        </div>
    </div>
</body>

</html>