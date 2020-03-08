INSERT INTO
   `activities`(field, new_value, old_value, data, description, entity_id, entity_type, consumer_id, action_by, source, type) 
   select
      t1.field,
      t1.new_value,
      t1.old_value,
      null as data,
      null as description,
      t2.entity_id,
      t2.entity_name as entity_type,
      t2.app_id as consumer_id,
      t3.email as action_by,
      null as source,
      null as type 
   from
      aunty.aunty_change as t1 
      left join
         aunty.aunty_action as t2 
         on t1.`action_id` = t2.id 
      join
         aunty.account_userprofile as t3 
         on t3.id = t2.profile_id