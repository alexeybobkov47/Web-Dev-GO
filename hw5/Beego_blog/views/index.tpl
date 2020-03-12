<!DOCTYPE html>

<html>

<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="../static/css/style.css">
</head>

<body>
  <div class="layout">
    <div class="header">
      <h1>{{.Blog.Name}}</h1>
      <h3>{{.Blog.Description}}</h3>
    </div>
    <div class="posts">
      {{range .Post}}
      <div class="post" id="{{.ID}}">
        <a href="post/{{.ID}}">
          <h5>{{.Header}}</h5>
        </a>
        <p>{{.Text}}</p>
        <a href="edit/{{.ID}}">Редактировать</a>
      </div>
      {{end}}
    </div>

  </div>
  <a href="newpost">Новый пост</a>

  <script src="../static/js/reload.min.js"></script>
</body>

</html>