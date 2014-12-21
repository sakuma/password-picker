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
    <li>
    <a href="{{password.URL}}">{{ password.Title }}</a>
    </li>
    {% endfor %}</ul>
    <hr>
    <a href="/passwords/new">New Password</a>
  </div>
</body>
</html>
