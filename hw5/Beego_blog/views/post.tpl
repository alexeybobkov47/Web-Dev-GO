<!DOCTYPE html>
<html lang="en">

<head>
    <link rel="stylesheet" href="../static/css/style.css">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>

<body>
    <div class="post" id="{{.ID}}">
        <h5>{{.Post.Header}}</h5>
        <p>{{.Post.Text}}</p>
    </div>
    <a href="/blog">На главную</a>
</body>

</html>