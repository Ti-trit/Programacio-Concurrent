with Ada.Numerics.Discrete_Random;
with Text_IO; use Text_IO;
package body def_monitor is

   
     --funció que genera un numero aleatori en el rang [MIN, MAX]
   function randomNumber(MIN, MAX:Integer) return Integer is
         type range_type is new Integer range MIN..MAX;
         package Rand_Int is new Ada.Numerics.Discrete_Random(range_type);
         use Rand_Int;
         Gen: Generator;
         num : range_type;
            begin
         reset(Gen);
         num:= random(gen);
         return Integer(num);
         
   end randomNumber;
   
   protected body monitor is 
      
      entry posarMel(id: in Id_Abella) when (pot<capacitatMax) is
         
         begin 
         Put_Line("L'abella " & id'Img & " ha posat una porció al pot");
         pot:=1+pot;
         
         if (pot = capacitatMax)then
            Put_Line("L'abella " & id'Img & " avisaré al ós");
            end if;
      
      end posarMel;
      
      
      
      
      entry consumirMel when (pot=capacitatMax) is
         begin 
         Put_Line("L'ós: menjaré tota la mel, yummy");
         pot:= 0;
         
         end consumirMel;
   
  
 end monitor;

end def_monitor;
