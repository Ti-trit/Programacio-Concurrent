with Text_IO;               use Text_IO;
with def_monitor;           use def_monitor;

procedure Main is

   m : monitor;


   task type  Abella is
   entry Start (id: in Id_Abella);
end Abella;


   task body Abella is
      My_id: Id_Abella;


      --presentar-se
   begin
      accept Start(id:Id_Abella) do
         My_id:= id;

      end Start;

      Put_Line("Hola, som l'abella" & My_id'Img );
      delay Duration (randomNumber (0, 5));

      for i in 1..REPETICIONS loop
         --posar mel
           m.posarMel(My_id);
           delay Duration (randomNumber (10, 20));
         end loop;


   end Abella;

   task type Os is
      entry Start;
   end Os;


   task body Os is

   begin
       accept Start do

         null;
      end Start;

      for i in 1..REPETICIONS loop
        delay Duration (randomNumber (2, 7));
         m.consumirMel;
      end loop;
      end Os;


   type threads_abelles is array (Integer range<>) of Abella;
   type threadOs is new Os;

   type abelles_access is access threads_abelles;
   type os_access is access threadOs;
   t1:  abelles_access;
   t2: os_access;



begin


         t1:= new threads_abelles(1..NUM_ABELLES);
         for i in 1..NUM_ABELLES loop
            t1(i).Start(Id_Abella(i));

         end loop;
      t2:= new threadOs;
   t2.Start;


end Main;
