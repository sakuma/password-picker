<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="UTF-8">
    <title>{{ password.Title }}</title>
    <link href="/assets/stylesheets/normalize.css" media="all" rel="stylesheet" />
    <link href="/assets/stylesheets/skeleton.css " media="all" rel="stylesheet" />
  </head>

  <body>
  <div class="container">
    <h1 class="title">「{{ password.Title }}」</h1>
    <form action="{{ password.URL }}" method="POST">
      <label for="title">Subject</label>
      <input id="title" name="title" class="u-full-width" type="text" value="{{ password.Title }}">
      <label for="body">Password</label>
      <input id="body" name="body" class="u-full-width" type="text" value="{{ password.Body }}">
      <input type="submit" value="Save">
    </form>
    <hr>
    <a href="/passwords">To Index</a>
    <a href="/passwords/{{ password.Id }}/delete" onclick="alert('Delete It. OK?');">削除</a>
    </div>
  </body>

</html>
