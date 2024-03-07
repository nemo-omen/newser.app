import { getHighlighter } from 'https://esm.sh/shiki@1.0.0'

async function loadTheme() {
  const themeResponse = await fetch('/static/theme/tokyo-night-jcis.json');
  console.log(themeResponse);
  if(!themeResponse.ok) {
    return;
  }
  
  let themeData;
  try {
    themeData = await themeResponse.json();
  } catch(err) {
    console.error(err);
    return;
  }
  
  const highlighter = await getHighlighter({
    langs: ['bash', 'html', 'scss', 'css', 'javascript', 'go', 'python', 'cpp', 'yaml', 'json', 'xml'],
    themes: []
  })
  
  await highlighter.loadTheme(themeData);
  
  const codeBlocks = document.querySelectorAll('pre>code');
  
  
  for(const block of codeBlocks) {
    const innerCode = block.innerText;
    const langDeclaration = block.className;
    const split = langDeclaration.split('-');
    const lang = split[split.length - 1];
    const parent = block.parentElement;
    
    parent.outerHTML = highlighter.codeToHtml(innerCode, {
      lang: lang,
      theme: 'tokyo-night-jcis'
    })
  }
}

loadTheme();