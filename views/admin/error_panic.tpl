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
                <h1 class="text-center">Panic Error</h1>
                <h2 class="text-center">{{.flash.Error}}</h2>
            </div>
        </div>
    </div>
</body>

</html>