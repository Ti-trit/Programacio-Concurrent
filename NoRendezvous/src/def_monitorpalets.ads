package def_monitorpalets is

   NUMFILOSOFS : constant Integer := 5;
   type index is new Integer range 0..NUMFILOSOFS - 1;
   type tpalets is array (index) of integer;
   function esq(i: index) return index;

   function dre(i: index) return index;


   protected type MonitorPalets is
      entry Agafa (index);
      procedure Deixa (Idx : in index);
      procedure Inicialitzar;
      procedure Acaba;


   private

   palets: tpalets;

   end MonitorPalets;

end def_monitorpalets;
