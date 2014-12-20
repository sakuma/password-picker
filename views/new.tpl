<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>パスワード登録</title>
  </head>
  <body>
    <div id="content">
      <form action="/passwords" method="POST">
      <h1 class="title">パスワード登録</h1>
      <p>
        Title:
        <input id="title" name="title" cols="80" rows="20" type="text" value="{{ password.Title }}" />
      </p>
      <p>
        パスワード:
        <input id="body" name="body" cols="80" rows="20" type="text" value="{{ password.Body }}" />
      </p>
      <p>
        <input type="submit" value="登録">
      </p>
      </form>
    </div>
  </body>
</html>
