with Ada.Numerics.Discrete_Random;
with Text_IO; use Text_IO;

package body def_monitor is
   
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

    
      
       function get_num_clients_actius return Integer is
   begin
      return MaquinaMonitor.num_clients_actius;
      end get_num_clients_actius;
      
     -- function get_num_reposadors_acabats return Integer is
     -- begin 
       --  return MaquinaMonitor.num_reposadors_acabats;
      --end get_num_reposadors_acabats;
      
      
      function get_num_refrescs return Integer is 
      begin
         return MaquinaMonitor.NUM_REFRESCS;
         end get_num_refrescs;
     
      
      entry actualitzar_num_clients_actius when (num_clients_actius>0) is 
         begin
         MaquinaMonitor.num_clients_actius:= MaquinaMonitor.num_clients_actius-1;
         end actualitzar_num_clients_actius;
     
      
      procedure initialize_values is
          
      begin
         CLIENTS := randomNumber(MIN, MAX);
      --   CLIENTS := 0;
         REPOSADORS := randomNumber(MIN, MAX);
      -- REPOSADORS:=10;
       num_clients_actius:= CLIENTS;
       num_reposadors_acabats:= 0;  
        Put_Line("Simulació amb " & CLIENTS'img & " clients i " & REPOSADORS'img & " reposadors");
         Put_Line("********** La màquina està preparada");

      end initialize_values;

   -- funció que selecciona un numero aleatori en el range [MIN,MAX]
     
     -- entry acabar (id: in Reposador_Id) when True is
       --  begin 
         --Put_Line("El reposador " & id'img & " acaba i se'n va >>>>>>>>>>");
         --num_reposadors_acabats:= num_reposadors_acabats+1;
       --  Put_Line("Numero de clients actius ara es " & get_num_clients_actius'img);
         --end acabar;
         
      
      
    entry posarRefrescs(id: in Reposador_Id) when True  is
      
         quantitat: Integer:= Max_refrescs-get_num_refrescs;
      begin 
         
        -- posant:= true;
        --Els clients encara no han arribat pero no són nuls
        if (NUM_REFRESCS<Max_refrescs) then 
         if num_clients_actius=0 then 
            Put_Line("++++++++++ El reposador " & id'Img & " diu: No hi ha clients m'en vaig");
        
         else
         Put_Line("++++++++++ El reposador " & id'Img & " reposa " & quantitat'Img & " refrescs, ara n'hi ha " & Max_refrescs'Img);
         NUM_REFRESCS:= Max_refrescs;--omplir la maquina fins al màxim
         end if;
         end if;
         
         
      end posarRefrescs;
      
      procedure release_reposador is 
         begin
         posant:= false;
      end release_reposador;
      
      procedure release_client is 
      begin
         consumint:= false;
         end release_client;

      
   entry consumir (nom: in nomClient; maxConsumacions: in Integer; cons_fetes: in out Integer)  when (NUM_REFRESCS>1)  is
      begin 
        -- consumint:=true;
           
         cons_fetes:= cons_fetes+1;
         NUM_REFRESCS:= NUM_REFRESCS-1;
     
         Put_Line("---------- " & To_String(nom) & " agafa el refresc " & cons_fetes'Img & "/" & maxConsumacions'Img & " a la màquina en queden " &NUM_REFRESCS'Img);
       
      
   end consumir;

    
     
     end MaquinaMonitor;

end def_monitor;
