<!DOCTYPE html>
<html lang="en">

<head>
    <link rel="stylesheet" href="../static/css/style.css">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>

<body>
    <h1>Изменить пост</h1>
    <form method="GET">
        <div>
            <label>header</label>
            <input type="text" name="header" value="{{.Post.Header}}">
        </div>
        <div>
            <label>text</label>
            <input type="text" name="text" value="{{.Post.Text}}">
        </div>
        <input type="submit" value="submit">
    </form>
    <form method="GET">
        <input type="hidden" name="id" value="{{.Post.ID}}">
        <input type="submit" name="delete" value="Удалить">
    </form>
    <a href="/blog">На главную</a>
</body>

</html>