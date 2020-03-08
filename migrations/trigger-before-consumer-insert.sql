DELIMITER //
CREATE TRIGGER `BeforeConsumerInsert`
BEFORE INSERT ON consumers 
FOR EACH ROW
BEGIN
    IF NEW.token IS NULL THEN
       SET NEW.token = SUBSTR(SHA1(uuid()), 1, 32);
    END IF;
 
END //
DELIMITER ; 


drop trigger BeforeConsumerInsert
