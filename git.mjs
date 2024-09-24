import { execSync } from "node:child_process";
import readline from "node:readline";

// Function to run a command and handle errors
const runCommand = (command) => {
	try {
		execSync(command, { stdio: "inherit" });
	} catch (error) {
		console.error(`Error running command: ${command}`);
		process.exit(1);
	}
};

// Create readline interface for input
const rl = readline.createInterface({
	input: process.stdin,
	output: process.stdout,
});

// Add changes to staging
runCommand('git add . -- ":!.env"');

// Prompt for commit message
rl.question("Enter commit message: ", (msg) => {
	runCommand(`git commit -m "${msg}"`);
	rl.close();
});
