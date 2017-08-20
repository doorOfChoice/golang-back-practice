<!DOCTYPE html>
<html>
    <head></head>
    <body>
        <h1>{{.flash.error}}</h1>
        <h1>{{.flash.success}}</h1>
        <form action="/account/register" method="POST">
            <input type="text" name="username">
            <input type="text" name="pass-one">
            <input type="text" name="pass-two">
            <input type="submit">
        </form>
    </body>
</html>