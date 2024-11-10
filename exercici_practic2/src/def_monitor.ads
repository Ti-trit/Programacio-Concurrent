with Ada.Numerics.Discrete_Random;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;


package def_monitor is
   MIN: constant Integer:= 0;
   MAX: constant Integer:= 10;
   Max_refrescs: constant Integer := 10; --maxim numeor de refrescs

   function randomNumber(MIN, MAX:Integer) return Integer;

  CLIENTS:Integer;
  REPOSADORS: Integer;
   
   type Reposador_Id is new Integer range 1..MAX; --identificador numeric 
   type nomClient is new Unbounded_String;
   type nomsClients is array (1..MAX) of nomClient;--array de noms pels clients


   

   protected type MaquinaMonitor is
      entry posarRefrescs(id: in Reposador_Id); --un reposador omple la maquina
      entry consumir (nom: in nomClient; maxConsumacions: in Integer; cons_fetes: in out Integer);

      procedure release_reposador;
      procedure release_client;
      procedure initialize_values;
      function get_num_clients_actius return Integer;
      function get_num_refrescs return Integer;
      entry actualitzar_num_clients_actius;
      --procedure actualitzar_num_clients_actius;
      --entry acabar (id: in Reposador_Id) ;
      --procedure acabar (id: in Reposador_Id) ;
      --function get_num_reposadors_acabats return Integer;
   
   private
      
      posant: boolean:= false;
      consumint: boolean:= false;
      num_clients_actius: Integer := CLIENTS;
      num_reposadors_acabats:Integer:= 0;
      NUM_REFRESCS: Integer:= 0;

      
   
   end MaquinaMonitor;

end def_monitor;
