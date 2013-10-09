<?php
  require_once 'Page.php';
  if (post('password') == 'nobody ranks 14d') {
		$_SESSION['editable'] = true;
		if (post('login and query') && post('query')) {
			$db = new mysqli(
		    'cesoid.com', 'cesoida_wmgc', 'rlxn^PwbTob8', 'cesoida_wmgc'
		  );
		  if (!$db->query(post('query'))) { print $db->error; }
		} else {
			header("Location: http://{$_SERVER['HTTP_HOST']}/wmgc/players.php");
			exit;
		}
	}
?>
<html><body><form method='post'>
	<div>
		<label for='password'>Login</label>
		<input id='password' type='password' name='password' />
		<input type='submit' name='submit_button' value='login' />
	</div>
	<div>
		<label for='query'>Query</label>
		<textarea id='query' name='query' rows='6' cols='80'></textarea>
		<input type='submit' name='submit_button' value='login and query' />
	</div>
</form></body></html>