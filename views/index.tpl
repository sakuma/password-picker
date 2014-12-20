 <!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Password Picker</title>
</head>
<body>
  <div id="content">
    <h1 class="title">Password Picker</h1>
    <ul>
    {% for password in passwords %}
    <li>{{ password.Title }} </li>
    {% endfor %}</ul>
    <hr>
    <a href="/passwords/new">パスワード登録</a>
  </div>
</body>
</html>
