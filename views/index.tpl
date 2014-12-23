 <!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <title>Password Picker</title>
  <link href="/assets/stylesheets/normalize.css" media="all" rel="stylesheet" />
  <link href="/assets/stylesheets/skeleton.css " media="all" rel="stylesheet" />
</head>
<body>
  <div class="container">
    <h1 class="title">Password Picker</h1>
    <ul>
    {% for password in passwords %}
    <li>
    <a href="{{password.URL}}">{{ password.Title }}</a>
    </li>
    {% endfor %}</ul>
    <hr>
    <a href="/passwords/new">New Password</a>
  </div>
</body>
</html>
