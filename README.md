﻿# gs1_through_fmcg_data_syncer

The purpose of this script is to send product data from SAP to FMCG who then sends it to GS1. 
FMCG is nescessairy because they keep themselves updated on the industry and they have a REST API, whereas GS1 expect data in XML format.

The script listen for the SAP field: 
  - `U_CCF_Sync_GS1: 'Y'`

The scripts send errors through the designated teams channel and also posts the response from FMCG and GS1 back into SAP.
