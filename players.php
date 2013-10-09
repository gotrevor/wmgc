<?php require_once 'PlayersPage.php' ?>
<html>
  <head>
    <title>Players -- Western Mass Go Club</title>
    <meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1">
    <style>
      table.players td, table.players th {
        padding: 4px; border: 2px solid black; text-align: center;
      }
      table.players {
        float: left;
        margin-right: 20px;
        width: 392px;
        border: 2px solid black;
        border-collapse: collapse;
      }
      table.editable { float: none; }
    </style>
  </head>
  <body>
    <?php if ($_SESSION['editable']) { ?>
      <form method='post' style='text-align: right;'>
        <input type='submit' name='submit_button' value='logout' />
      </form>
    <?php } ?>
    <div style='text-align: center;'>
      <img src="westernmass.gif" alt="Western Massachusetts Go Club" />
    </div>
    <div style='text-align: center;'><ul><?= Page::nav_items(array(
      array('index', 'Western Mass Go Club'),
      array('wmgt', 'Tournaments'),
      array('', 'Players'),
      array('directions', 'Directions'),
      array('resources', 'Go Resources'),
    )) ?></ul></div>
    <?= PlayersPage::content() ?>
    <br clear='all' />
    <div style='text-align: right;'><a href='login.php'>&pi;</a></div>
  </body>
</html>