package def_monitor is

   NUM_NANS: Integer:= 7;
   NUM_MENJADES: Integer:= 2;
   
   function randomNumber(MIN, MAX:Integer) return Integer;
   type Nan_Id Is new Integer range 1..NUM_NANS;
   
   protected type monitor is
      entry seure(id: in Nan_Id);
    --  procedure passejar();
   --   procedure anarLaMina(id: in Nan_Id);
      entry dormir (id:in Nan_Id);
      entry demanarMenjar(id: in Nan_Id);
      
   private
      NUM_CADIRES: Integer:= 4;
      a_menjar:Integer:= 0;
      a_dormir: Integer:= 0;
      menjar_preparat: Boolean:= False;
end def_monitor;
