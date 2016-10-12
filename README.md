# tvinfo-server
This project is being designed as an open version of sites such as TheTVDB and TVRage for storing information about television shows. It is being written in go.

The planned design is likely a database style setup with a TV Series->Season->Episode Hierarchy. I would also like a macro series listing if a series uses it (Aka a series that has an overall title even though each "season" may have a different title)
With the Episode info containing:
	Title
	Episode Number
	Description
	Release Date
	Production code
	Guest Stars
	Director
	Writer
	Absolute Episode Number (Series Wide Episode Count with Absolute Positioning)
	DVD/Blu-Ray info (Each option containing the follow)
		DVD/BD Season
		DVD/BD Episode
		DVD/BD Disc
		DVD/BD Title/Track Number

Shows should be able to be locked at the macro, series, season, episode level by admins (much like is done with wikipedia articles and such when things get into edit wars)
Art to be able to be associated with each level of the Hierarchy.

An API setup to allow other sites, programs, etc to scrape/pull the information to be used by others.
