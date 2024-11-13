with Ada.Numerics.Discrete_Random;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;


package def_monitor is
   MIN: constant Integer:= 0;
   MAX: constant Integer:= 10;
   Max_refrescs: constant Integer := 10; --maxim numero de refrescs

   function randomNumber(MIN, MAX:Integer) return Integer;

  CLIENTS:Integer;
  REPOSADORS: Integer;
   
   type Reposador_Id is new Integer range 1..MAX; --identificador numeric 
   type nomClient is new Unbounded_String;
   type nomsClients is array (1..MAX) of nomClient;--array de noms pels clients


  
   protected type MaquinaMonitor is
      entry posarRefrescs(id: in Reposador_Id); 
      entry consumir (nom: in nomClient; maxConsumacions: in Integer; cons_fetes: in out Integer);
      entry actualitzar_num_clients_actius;

      procedure initialize_values;
      function get_num_clients_actius return Integer;

   
   private
     
      num_clients_actius: Integer := CLIENTS;
      NUM_REFRESCS: Integer:= 0;

  
   
   end MaquinaMonitor;

end def_monitor;
