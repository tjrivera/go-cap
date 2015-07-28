/*
   Package redcap provides a REDCap client implementation focusing
   around the concept of a REDCap project.

   Most users will interact with REDCap in the context of a project:

           project := redcap.NewRedcapProject("http://redcap.myinstitution.org", "<redcap-token>", true)

   Retrieve a prject's metadata and forms:

           fields := project.GetMetadata()
           ...
           forms := project.GetForms()
           ...


*/
package redcap
