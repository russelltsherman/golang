{{ define "new" }}
<!DOCTYPE html>
<html>
<head></head>
<body>
<form action="new" method="post">
  <input type="text" name="title" placeholder="Title"><br>
  Content<br>
  <textarea type="text" name="content"></textarea><br>
  <input type="radio" name="content-type" value="page" checked>Page
  <input type="radio" name="content-type" value="post">Post
  <br>
  <input type="submit" value="Submit">
</form>
</body>
</html>
{{ end }}
