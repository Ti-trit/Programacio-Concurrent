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
   

    protected body MaquinaMonitor is 

      --getter de l'atribut private num_clients_actius
       function get_num_clients_actius return Integer is
   begin
      return MaquinaMonitor.num_clients_actius;
      end get_num_clients_actius;
      
      entry actualitzar_num_clients_actius when (num_clients_actius>0) is 
         begin
         MaquinaMonitor.num_clients_actius:= MaquinaMonitor.num_clients_actius-1;
         end actualitzar_num_clients_actius;
     
      
      procedure initialize_values is
          
      begin
         CLIENTS := randomNumber(MIN, MAX);
         REPOSADORS := randomNumber(MIN, MAX);
         num_clients_actius:= CLIENTS;
         Put_Line("Simulació amb " & CLIENTS'img & " clients i " & REPOSADORS'img & " reposadors");
         Put_Line("********** La màquina està preparada");

      end initialize_values;

   
    entry posarRefrescs(id: in Reposador_Id) when True  is
      
         quantitat: Integer:= Max_refrescs-NUM_REFRESCS;
      begin 
         
         --Si la màquina no está ple i encara hi ha clients actius
          
        if (NUM_REFRESCS<Max_refrescs and num_clients_actius>0) then 
        
       Put_Line("++++++++++ El reposador " & id'Img & " reposa " & quantitat'Img & " refrescs, ara n'hi ha " & Max_refrescs'Img);
         NUM_REFRESCS:= Max_refrescs;--omplir la maquina fins al màxim
         end if;
         
         
      end posarRefrescs;
      
      
   entry consumir (nom: in nomClient; maxConsumacions: in Integer; cons_fetes: in out Integer)  when (NUM_REFRESCS>0)  is
      begin 
       
         cons_fetes:= cons_fetes+1;
         NUM_REFRESCS:= NUM_REFRESCS-1;
     
         Put_Line("---------- " & To_String(nom) & " agafa el refresc " & cons_fetes'Img & "/" & maxConsumacions'Img & " a la màquina en queden " &NUM_REFRESCS'Img);
       
      
   end consumir;

   
     end MaquinaMonitor;

end def_monitor;
