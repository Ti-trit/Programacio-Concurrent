with Text_IO;               use Text_IO;
with def_monitor;           use def_monitor;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;

procedure Maquinarefrescs is
   monitor_maquina : MaquinaMonitor;

   -- Funció que converteix un string unbounded_string

   function tub (Source : String) return Unbounded_String renames
     Ada.Strings.Unbounded.To_Unbounded_String;

   -----
   -- Especificació de la tasca de client
   -----
   task type Client is
      entry Start (nom : in nomClient);
   end Client;

   -----
   -- Cos de la tasca de client
   -----

   task body Client is
      my_name         : nomClient; --identificador del client
      maxConsumacions : Integer := randomNumber (MIN, MAX); --numero de consumicions que vol fer
      cons_fetes      : Integer := 0;--numero de consumicions fetes fins ara
   begin
      accept Start (nom : in nomClient) do
         my_name    := nom;
         cons_fetes := 0;
      end Start;
            delay Duration (randomNumber (0, 5));

      Put_Line
        (To_String (my_name) & " diu: Hola, avui faré " & maxConsumacions'Img &
         " consumicions");

      if REPOSADORS = 0 then
         Put_Line(To_String (my_name) &" diu : No hi ha reposadors a la màquina, m'en vaig");
         monitor_maquina.actualitzar_num_clients_actius; --reduir numero de clients actius amb una unitat
         Put_Line(To_String (my_name) & " acaba i se'n va, queden " & monitor_maquina.get_num_clients_actius'Img & " clients >>>>>>>>>>");

      else
         -- hi ha reposadors a la màquina
      -- un bucle per el total de consumacions
         for i in 1 .. maxConsumacions loop
             delay Duration (randomNumber (1, 5));
            monitor_maquina.consumir (my_name, maxConsumacions, cons_fetes);

         end loop;
         monitor_maquina.actualitzar_num_clients_actius; --reduir numero de clients actius
         Put_Line(To_String (my_name) & " acaba i se'n va, queden " & monitor_maquina.get_num_clients_actius'Img & " clients >>>>>>>>>>");
      end if;

   end Client;

   type nomsClientsArray is array (1 .. MAX) of nomClient;

   nomsClients : nomsClientsArray :=
     (nomClient (tub ("Aina")), nomClient (tub ("Bel")),
      nomClient (tub ("Albert")), nomClient (tub ("Toni")),
      nomClient (tub ("Miguel")), nomClient (tub ("Clara")),
      nomClient (tub ("Estela")), nomClient (tub ("Sofia")),
      nomClient (tub ("Pere")), nomClient (tub ("Laura")));


   task type Reposador is
      entry Start (Idx : in Reposador_Id);
   end Reposador;

   task body Reposador is
      My_Idx : Reposador_Id; --identificador numeric del reposador

   begin

      accept Start (Idx : in Reposador_Id) do
         My_Idx := Idx;

      end Start;
      delay Duration(randomNumber(1,4));--retard de simulació

      Put_Line ("El reposador " & My_Idx'Img & " comença a treballar");
      delay Duration(randomNumber(0,3));

      if CLIENTS = 0 then
         Put_Line("++++++++++ El reposador " & My_Idx'Img & " diu: No hi ha clients m'en vaig");
      else
          --mentre hi hagi clients van posant refrescs a la maquina
      while monitor_maquina.get_num_clients_actius > 0 loop
         delay Duration(randomNumber(1,5));
         monitor_maquina.posarRefrescs (My_Idx);

      end loop;
   end if;

      delay Duration (randomNumber (1, 3));
      -- Si tots els clients han fet les seves consumacions, han d'anar acabant
      Put_Line ("El reposador " & My_Idx'Img & " acaba i se'n va >>>>>>>>>>");

   end Reposador;

   type threadsClients_Array is array (Integer range <>) of Client;
   type threadsReposadors_Array is array (Integer range <>) of Reposador;

-- Declarem tipus d'accés (punters) a aquests arrays
   type threadsClients_Access is access threadsClients_Array;
   type threadsReposadors_Access is access threadsReposadors_Array;

   -- Instancies dels arrays com punters
   t1 : threadsClients_Access;
   t2 : threadsReposadors_Access;

begin
-- inicialitzar els valos de CLIENS I REPOSADORS i inicialitzar la màquina
   monitor_maquina.initialize_values;
   --  començem les tasques dels clients i dels reposadors

   t1 := new threadsClients_Array (1 .. CLIENTS);
   t2 := new threadsReposadors_Array (1 .. REPOSADORS);

   for Idx in 1 .. REPOSADORS loop

      t2 (Idx).Start (Reposador_Id (Idx));

   end loop;

   for i in 1 .. CLIENTS loop

      t1 (i).Start (nomsClients (i));
   end loop;

end Maquinarefrescs;
