package def_monitor is

   NUM_ABELLES:constant Integer:= 10;
   REPETICIONS:constant Integer:=10;
   
  function randomNumber(MIN, MAX:Integer) return Integer;
   type  Id_Abella is new Integer range 1..NUM_ABELLES;
   
   
   protected type monitor is 
   entry posarMel(id:in Id_Abella);
      entry consumirMel;
      
   
     
private
   consumint:Boolean:= False;
   pot: Integer:= 0;
   capacitatMax:Integer:= 10;
   
      end monitor;

end def_monitor;
