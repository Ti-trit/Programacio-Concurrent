with Text_IO;               use Text_IO;
with def_monitor;           use def_monitor;

procedure Main is

   --instancia del monitor
   monitor: monitor;

   task type Blancaneus is
      entry Start();
   end Blancaneus;

   task type Nan is
      entry Start(id: in Nan_Id);
   end Nan;

   task body Blancaneus is

   begin
      accept Start() do
      end Start();

      Put_Line("Hola, som Blancaneus.");
      for i in 1...repeticions loop

         Put_Line("Me'n vaig a passejar");

         Put_Line("Preparant el menjar");





   end Blancaneus;












      type threads_nans is array (Integer range<>) of Nan;










      begin

         threads:= new threads_nans (1..NUM_NANS);
         for i in 1..NUM_NANS loop
            t1(i).Start(Nan_Id(i));
         end loop;



end Main;
