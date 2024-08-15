// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/ed-henrique/ed-henrique.com/pages/components"

func CV() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col gap-3 pt-3 w-full\"><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">About</h2><div class=\"text-sm sm:text-base\"><p>Eduardo Henrique Freire Machado</p><p>Boa Vista, Roraima</p><p><a class=\"border-b border-ctp-sapphire text-ctp-sapphire\" href=\"mailto:edu.hen.fm@gmail.com\" referrerpolicy=\"no-referrer\" rel=\"noreferrer\" target=\"_blank\">edu.hen.fm@gmail.com</a></p><p><a class=\"border-b border-ctp-sapphire text-ctp-sapphire\" href=\"https://linkedin.com/in/ed-hfm\" referrerpolicy=\"no-referrer\" rel=\"noreferrer\" target=\"_blank\">in/ed-hfm</a></p><p><a class=\"border-b border-ctp-sapphire text-ctp-sapphire\" href=\"https://orcid.org/0009-0008-7340-976X\" referrerpolicy=\"no-referrer\" rel=\"noreferrer\" target=\"_blank\">orcid.org/0009-0008-7340-976X</a></p></div></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Work Experience</h2><div><p class=\"mb-3 italic\">DevOps</p><p class=\"text-sm sm:text-base\">Backend Tech Lead</p><p class=\"text-sm sm:text-base\">- Developed and maintained scalable web applications using Go</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Maranhão, Brazil | Apr/2023 - Present</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Government of the State of Roraima</p><p class=\"text-sm sm:text-base\">Head of the Projects and Solutions Division</p><p class=\"text-sm sm:text-base\">- Developed and maintained scalable web applications using Go, Templ, and AlpineJS</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Roraima, Brazil | Mar/2023 - Present</p></div></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Skills</h2><div class=\"space-y-3\"><div><p class=\"mb-3 italic\">Programming Languages</p><p class=\"text-sm sm:text-base\">- Go, JavaScript, TypeScript, Python, C, Lua</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">- Rust, Zig, Elixir, R</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Web Technologies</p><p class=\"text-sm sm:text-base\">- Svelte, Node.js, TailwindCSS, HTML, CSS</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Cloud Platforms</p><p class=\"text-sm sm:text-base\">- AWS, Google Cloud Platform</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">DevOps Tools</p><p class=\"text-sm sm:text-base\">- Git, Docker, Kubernetes, Bash, GitHub Actions</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">- TrueNAS</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Database Systems</p><p class=\"text-sm sm:text-base\">- SQLite, PostgreSQL, MongoDB</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Soft Skills</p><p class=\"text-sm sm:text-base\">- Leadership, Communication, Problem-solving, Team Collaboration</p></div></div></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Impressive Achievements</h2></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Leadership and Collaboration</h2></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Educational Background</h2><div><p class=\"mb-3 italic\">Bachelor's Degree in Computer Science</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Universidade Federal de Roraima</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Roraima, Brazil | 2020-2025</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Degree in Mathematics</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Faculdade Única</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Roraima, Brazil | 2024-2024</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Lato Sensu Post-Graduate's Degree in Computer Systems Analysis</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Faculdade Única</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Roraima, Brazil | 2024-2024</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Lato Sensu Post-Graduate's Degree in Software Development</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Faculdade Focus</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Roraima, Brazil | 2024-2024</p></div><div class=\"border-t border-ctp-surface2 max-w-8\"></div><div><p class=\"mb-3 italic\">Technologist in Digital Marketing</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Universidade Claretiano</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">São Paulo, Brazil | 2020-2021</p></div></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Continuous Learning and Certifications</h2></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Recognition and Awards</h2></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Diverse Projects</h2></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Volunteering and Leadership</h2><div><p class=\"mb-3 italic\">Operação Acolhida</p><p class=\"text-sm sm:text-base\">Volunteer Teacher</p><p class=\"text-sm sm:text-base\">- Taught refugees how to get into work in a foreign country improving their network reach and digital presence</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">Roraima, Brazil | 2022</p></div></section><section class=\"flex flex-col gap-3\"><h2 class=\"font-bold\">Publications</h2><div><p class=\"mb-3 italic\">Percepção de Alunos sobre a Eficiência do Jogo \"Enigma das Frações\" no Aprendizado de Números Racionais</p><p class=\"text-sm sm:text-base\">Writer and Reviewer</p><p class=\"text-sm sm:text-base\">- Reviewed an academic paper to turn it into a book chapter over the course of a month</p><p class=\"text-ctp-overlay0 text-sm sm:text-base\">São Paulo, Brazil | 2023</p></div></section></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = components.Base("cv").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
