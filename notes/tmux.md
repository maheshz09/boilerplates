#tmux

IMP:
	ctrl + B = perfix key 


######################################## session ########################################
-- check the exesting session running
# tmux ls

-- to attach to any 
# tmux attach (attach to latest/current/recent one)
# tmux attach -t 0 (attach with the index number -t means target)

-- detach from any
# inside tmux
= ctrl + B + D

-- create new tmux session 
# tmux new -s <name of the session>
OR 
# tmux -- it will create new session by itself

-- kill any session
# tmux kill-session (this will kill the latest session)
for specific session like 
# tmux kill-session -t <session-name>


##################################### window ###############################################

-- to create new session & window (default it crate the window)
# tmux new -t test
	we can see session and session specific info down there only
	
-- for vertical speration of terminal will use (inside session)
# ctrl + B + shift + %

-- for horizontal speration of the terminal
# ctrl + B + shift + "

NOW for split terminal (vertical/horizontal) how to switch between them NOTE: they highligh "green" by default
# ctrl + B + direction arrorws on keyboard(up, donw, left, right)

	---another method for the same----
	# ctrl + B + Q
		if we did this it will show us in index of every pain (horizontal/vertical) that you have created in single windows
	
	--- to swich to desired index of plain
	# ctrl + B + Q + <index number>
	

-- if we want to change the size of the pain 
# ctrl+B hold ctrl again & resize the pain with (up, donw, left, right) keys
	
	-- we can able to use template
	# ctrl + B + alt (1,2,3,4 keys)



----- how to create new window
# ctrl + B + C
	-- use specific windows 
	# ctrl + B + <index number listed below>
	

	-- to rename windows
	# ctrl + B + <
	
----- TO MANAGE ENTIRE SESSION WINDOWS EVERYTHING (LIST OF SESSION/WINDOWS)
# ctrl + B + W

	-- to kill specific pain from windows using same mode
	# ctrl + B + X (it will ask you want to kil or not)
	
	-- to delete all at once (window only)
	# ctrl + B + &
	
	

------- to kill all session/windows everything
# tmux kill-server

--- copy mode in tmux (before this we need to configre tmux.rc)
# ctrl + B + { (this will start your copy mode)
	use directional arrow from whare we have to copy
	# space key to start	
	