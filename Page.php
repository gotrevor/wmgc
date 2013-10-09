<?php
  ini_set('display_errors', '1');
  function post($field) {
    if (isset($_POST[$field])) { return $_POST[$field]; }
  }
  session_start();
  if (post('submit_button') == 'logout') {
    $_SESSION['editable'] = false;
  }
  class Page {
    function nav_items($pages) {
      $items = array();
      foreach ($pages as $page) { ob_start();
        list($page, $title) = $page;
        ?><li style='display: inline;'>
          <?php if ($page) { ?><a href="<?= $page ?>.html"><?=
            $title
          ?></a><?php } else
          { print $title; } ?>
        </li><?php
      $items[] = ob_get_clean(); }
      return implode(' | ', $items);
    }
  }
?>