<!DOCTYPE html>
<html>
    <head></head>
    <body>
        <h1>{{.flash.error}}</h1>
        <h1>{{.flash.success}}</h1>
        <form action="/account/login" method="POST">
            <input type="text" name="username">
            <input type="text" name="password">
            <input type="submit">
        </form>
    </body>
</html>