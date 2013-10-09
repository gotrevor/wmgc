<?php
	require_once 'Page.php';
	class PlayersPage extends Page {
		function content() { ob_start();
		  $db = new mysqli(
		    'cesoid.com', 'cesoida_wmgc', 'rlxn^PwbTob8', 'cesoida_wmgc'
		  );
      if ($names = post('name')) {
	      foreach (array('new' => 'insert', 'existing' => 'update') as $type => $query) {
	       	if (isset($names[$type])) {
		       	foreach ($names[$type] as $index => $name) {
		       		$sets = array();
		       	 	foreach (self::attributes() as $attribute) { $sets[] = "$attribute='".
		       	 		$db->real_escape_string($_POST[$attribute][$type][$index])
		       	 	."'"; }
		       	 	$db->query("$query players
		       	 	 	set ".implode(', ', $sets)."
		       	 	 	".($query == 'update' ?
		       	 	 		"where id='".$db->real_escape_string($index)."'" :
		       	 	 		''
		       	 	 	)
		       	 	);
		       	}
		      }
		    }
	    }
		  ?>
		    <p>
		      Click on the player's name to see their picture.  If you think your name
		      should be in this list, please e-mail
		      <a href="mailto:toddcesere@gmail.com">toddcesere@gmail.com</a>.
		    </p>
		    <h2>Players</h2>
		    <?= self::player_table($db->query('select * from players where current = 1 limit 13')) ?>
		    <?= self::player_table($db->query('select * from players where current = 1 limit 13, 13')) ?>
		    <br clear='all' />
		    <h2>Alumni</h2>
		  <?php
	    print self::player_table($db->query('select * from players where current = 0'));
	    return ob_get_clean();
		}
    function player_table($player_result) { ob_start();
      $editable = $_SESSION['editable'];
      $attributes = self::attributes();
      if ($editable) { ?><form method='post'>
        <input type='submit' name='submit_button' value='save' />
        <?php } ?>
          <table class='players <?= $editable ? 'editable' : '' ?>'>
            <tr><?php foreach (self::labels() as $column) {
              ?><th><?= $column ?></th><?php
            } ?></tr>
            <?php
            	$count = 0;
	            while ($player = $player_result->fetch_assoc()) {
	            	++$count;
	              extract($player);
	              ?><tr><?php
	               	if ($editable) {
		               	foreach ($attributes as $attribute) {
			                $attribute = strtolower($attribute);
			                ?><td><input
			                  name='<?= $attribute.($id ? "[existing][$id]" : "[new][$count]") ?>'
			                  value='<?= htmlspecialchars($$attribute) ?>' 
			                /><?= $attribute == 'photo' ? '.jpg' : '' ?></td><?php
			              }
		              } else { ?>
		                <td><?php
		                  if ($photo) { ?>
		                    <a href="playerspics/<?= $photo ?>.jpg" target="_blank"><?=
		                      $name
		                    ?></a>
		                  <?php } else {
		                   	print $name;
		                  }
		                ?></td>
		                <?php for ($i = 1; $i < count($attributes); ++$i) { ?>
		                  <td><?= $$attributes[$i] ?></td>
		                <?php }
	              	}
	              ?></tr><?php
	            }
	          ?>
          </table>
        <?php if ($editable) { ?>
      </form><?php }
    return ob_get_clean(); }
		function attributes() { return array_map('strtolower', self::labels()); }
	  function labels() {
	    $attributes = array('Name', 'Rank', 'IGS', 'KGS');
	    if ($_SESSION['editable']) {
	    	array_unshift($attributes, 'Photo');
	    	array_push($attributes, 'Current');
	    }
	    return $attributes;
	  }
	}
?>