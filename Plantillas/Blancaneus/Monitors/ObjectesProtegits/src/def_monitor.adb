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
      
      
      
      entry seure(id: in Nan_Id)  when (NUM_CADIRES>0)is
         NUM_CADIRES:= NUM_CADIRES-1;
         Put_Line("El nan " & id'Img & "ha trobat una cadira buida");
              
      end seure;
      
      entry dormir (id:in Nan_Id, numMenjades:Integer) when (True) is
         if (numMenjades = NUM_MENJADES) then
            Put_Line("---->> El nan " & id'Img & "ha anat a dormir");
            monitor.a_dormir := monitor.a_dormir+1;
      
     
      
      end monitor;

ççend def_monitor; 
