with Text_IO;
use Text_IO;
with def_monitor;
use def_monitor;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;

-- Funció que converteix un string a una de unbounded_string


procedure Maquinarefrescs is
   monitor_maquina: MaquinaMonitor;


   function tub(Source : String) return unbounded_string renames ada.strings.unbounded.to_unbounded_string;

 --  type Reposador_Id is new Integer range 0..10-1; --identificador numeric

   -----
  	-- Especificació de la tasca de client
  	-----
   task type Client is
      entry Start(nom: in nomClient);
   end Client;

   -----
  	-- Cos de la tasca de client
   -----

   --type nomClient is new Unbounded_String;
   task body Client is
      my_name: nomClient;
      maxConsumacions: Integer := randomNumber(MIN,MAX);
      cons_fetes: Integer:= 0;
   begin
      accept Start (nom: in nomClient) do
         my_name := nom;
         cons_fetes:= 0;
      end Start;

      Put_Line(To_String(my_name) & " diu: Hola, avui faré " & maxConsumacions'Img & " consumicions");
      delay(1.0);


      -- un bucle per el total de consumacions
      if REPOSADORS=0 then
         Put_Line( To_String(my_name) & " diu : No hi ha reposadors a la màquina, m'en vaig");
          monitor_maquina.actualitzar_num_clients_actius; --reduir numero de clients actius
         Put_Line( To_String(my_name) & " acaba i se'n va, queden " & monitor_maquina.get_num_clients_actius'img & " clients >>>>>>>>>>");

      else
      -- hi ha reposadors a la màquina
      --
      for i in 1..maxConsumacions loop
         monitor_maquina.consumir(my_name, maxConsumacions, cons_fetes);
       --  monitor_maquina.release_client;

        delay Duration(2.0);

        end loop;
         monitor_maquina.actualitzar_num_clients_actius; --reduir numero de clients actius
         Put_Line( To_String(my_name) & " acaba i se'n va, queden " & monitor_maquina.get_num_clients_actius'Img & " clients >>>>>>>>>>");
      end if;

   end Client;



   type nomsClientsArray is array (1..MAX) of nomClient;

   nomsClients: nomsClientsArray := (nomClient(tub("Aina")), nomClient(tub("Bel")), nomClient(tub("Albert")), nomClient(tub("Toni")),nomClient(tub("Miguel")),
                                    nomClient(tub("Clara")),nomClient(tub( "Estela")), nomClient(tub("Sofia")), nomClient(tub("Pere")), nomClient(tub("Laura")));

procedure posar_refrescs (Idx: in Reposador_Id) is
      begin
      monitor_maquina.posarRefrescs(Idx);
    --  monitor_maquina.release_reposador;
      delay Duration (1.22);

   end posar_refrescs;


   task type Reposador is
      entry Start(Idx: in Reposador_Id);
   end Reposador;

   task body Reposador is
      My_Idx: Reposador_Id;

   begin

      accept Start (Idx: in Reposador_Id) do
         My_Idx:= Idx;

      end Start;
      Put_Line("El reposador " & My_Idx'img & " comença a treballar");
      --mentre hi hagi clients van posant refrescs a la maquina

      while monitor_maquina.get_num_clients_actius > 0 loop
        -- delay(1.0);
        monitor_maquina.posarRefrescs(My_Idx);

         end loop;

        delay(1.2);
      -- Si tots els clients han fet les seves consumacions, han d'anar acabant
      Put_Line("El reposador " & My_Idx'img & " acaba i se'n va >>>>>>>>>>");



   end Reposador;



   -- Crear els processos concurrents de clients i reposadors

  --type threadsClients is array (1..MAX) of Client;
  --t1: threadsClients;

  --type threadsReposadors is array (1..MAX) of Reposador;
  --t2: threadsReposadors;

  type threadsClients_Array is array (Integer range <>) of Client;
  type threadsReposadors_Array is array (Integer range <>) of Reposador;

   -- Declaramos tipos de acceso (punteros) a estos arreglos
  type threadsClients_Access is access threadsClients_Array;
  type threadsReposadors_Access is access threadsReposadors_Array;

   -- Instancias de los arreglos como punteros
   t1 : threadsClients_Access;
   t2 : threadsReposadors_Access;

begin


   monitor_maquina.initialize_values; -- inicialitzar els valos de CLIENS I REPOSADORS i inicialitzar la màquina
   --  començem les tasques dels clients i dels reposadors

 t1 := new threadsClients_Array(1 .. CLIENTS );
   t2 := new threadsReposadors_Array(1 .. REPOSADORS);

  for Idx in 1..REPOSADORS loop

   t2(Idx).Start(Reposador_Id(Idx));

   end loop;

  -- Put_Line("array thread de RESPODAORS" );

    for i in 1..CLIENTS loop


         t1(i).Start(nomsClients(i));
   end loop;


end Maquinarefrescs;

